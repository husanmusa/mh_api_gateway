package client

import (
	"github.com/husanmusa/mh_api_gateway/config"

	pb "github.com/husanmusa/mh_api_gateway/genproto/appointment_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	AppointmentService() pb.AppointmentServiceClient
	DoctorService() pb.DoctorServiceClient
	PatientService() pb.PatientServiceClient
}

type grpcClients struct {
	appointmentService pb.AppointmentServiceClient
	doctorService      pb.DoctorServiceClient
	patientService     pb.PatientServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connAppointService, err := grpc.Dial(
		cfg.AppointmentServiceHost+cfg.AppointmentGRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		appointmentService: pb.NewAppointmentServiceClient(connAppointService),
		doctorService:      pb.NewDoctorServiceClient(connAppointService),
		patientService:     pb.NewPatientServiceClient(connAppointService),
	}, nil
}
func (g *grpcClients) AppointmentService() pb.AppointmentServiceClient {
	return g.appointmentService
}
func (g *grpcClients) DoctorService() pb.DoctorServiceClient {
	return g.doctorService
}
func (g *grpcClients) PatientService() pb.PatientServiceClient {
	return g.patientService
}
