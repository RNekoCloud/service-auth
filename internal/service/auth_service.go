package service

import (
	"fmt"
	"log"

	"github.com/cvzamannow/service-auth/internal/entity"
	"github.com/cvzamannow/service-auth/internal/helper/email"
	"github.com/cvzamannow/service-auth/internal/helper/token"
	pb "github.com/cvzamannow/service-auth/internal/proto"
	"github.com/cvzamannow/service-auth/internal/repository"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	Repository repository.AuthRepository
}

func (srv *Server) SignUp(req *pb.SignUpRequest, stream pb.AuthService_SignUpServer) error {
	isValid := email.IsEmail(req.Email)

	if !isValid {
		return stream.Send(&pb.SignUpResponse{
			Message: "Email tidak valid!",
			Status:  "failure",
		})
	}

	sendMessage([]byte(req.Email))

	id := uuid.New().String()
	pass := []byte(req.Password)
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		return err
	}

	_, errValidate := srv.Repository.CheckUser(req.Email)
	if errValidate == nil {
		return stream.Send(&pb.SignUpResponse{
			Message: "Email sudah terdaftar!",
			Status:  "failure",
		})
	}

	_ , err = srv.Repository.SaveUser(&entity.UserEntity{
		Id:       id,
		Email:    req.Email,
		Password: string(hash),
		Role:     0,
	})
	if err != nil {
		return err
	}

	return nil
}

func (*Server) TestMessage(req *pb.TestRequest, stream pb.AuthService_TestMessageServer) error {
	return stream.Send(&pb.TestResponse{
		Message: "Foo",
	})
}

func (srv *Server) SignIn(req *pb.SignInRequest, stream pb.AuthService_SignInServer) error {
	isValid := email.IsEmail(req.Email)
	if !isValid {
		return stream.Send(&pb.SignInResponse{
			Message: "Format email tidak valid!",
			Token:   "none",
		})
	}

	user, err := srv.Repository.CheckUser(req.Email)
	if err != nil {
		return stream.Send(&pb.SignInResponse{
			Message: "Email tidak terdaftar! Silahkan register terlebih dahulu",
			Token:   "none",
		})
	}

	hashed := []byte(user.Password)
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(req.Password)); err != nil {
		return stream.Send(&pb.SignInResponse{
			Message: "Password salah!",
			Token:   "none",
		})
	}

	token, err := token.EncodeToken(user.Id, user.Email)
	if err != nil {
		return err
	}

	return stream.Send(&pb.SignInResponse{
		Message: "Berhasil login!",
		Token:   token,
	})
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func sendMessage(body []byte) error {
	conn, err := amqp.Dial("amqp://guest:guest@ec2-3-112-2-190.ap-northeast-1.compute.amazonaws.com:5672/") // Ganti dengan koneksi RabbitMQ Anda
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // Nama antrian
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		return err
	}

	fmt.Printf(" [x] Sent %s\n", body)
	return nil
}
