package mask

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// START OMIT
func UnaryServerInterceptor(cs ...codes.Code) grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		_ *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		resp, err := handler(ctx, req) // HL
		if err != nil {                // HL
			errCode := status.Code(err) // HL
			for _, c := range cs {      // HL
				if errCode == c { // HL
					err = status.Error(codes.Internal, "Internal server error") // HL
				} // HL
			} // HL
			return nil, err // HL
		} // HL
		return resp, nil // HL

	}
}

// END OMIT
