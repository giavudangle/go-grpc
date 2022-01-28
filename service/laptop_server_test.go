package service_test

import (
	"context"
	"testing"

	"github.com/giavudangle/go-grpc/pb"
	"github.com/giavudangle/go-grpc/sample"
	"github.com/giavudangle/go-grpc/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopWithNoId := sample.NewLaptop()
	laptopWithNoId.Id = ""

	laptopWithInvalidId := sample.NewLaptop()
	laptopWithInvalidId.Id = "invalid-uuid-here"

	laptopWithDuplicateId := sample.NewLaptop()
	storeWithDuplicateId := service.NewInMemoryLaptopStore()
	err := storeWithDuplicateId.Save(laptopWithDuplicateId)
	require.Nil(t, err)

	test_cases := []struct {
		name   string
		laptop *pb.Laptop
		store  service.LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_with_no_id",
			laptop: laptopWithNoId,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failure_with_invalid_id",
			laptop: laptopWithInvalidId,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failure_with_duplicate_id",
			laptop: laptopWithDuplicateId,
			store:  storeWithDuplicateId,
			code:   codes.AlreadyExists,
		},
	}

	for i := range test_cases {
		tc := test_cases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}
			server := service.NewLaptopServer(tc.store)

			res, err := server.CreateLaptop(context.Background(), req)

			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)

				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, res.Id)
				} else {
					require.NoError(t, err)
					require.Nil(t, res)
					st, ok := status.FromError(err)

					require.True(t, ok)
					require.Equal(t, tc.code, st.Code())

				}
			}
		})
	}
}
