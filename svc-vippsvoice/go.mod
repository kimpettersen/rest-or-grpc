module github.com/kimpettersen/go-api-talk/svc-vippsvoice

go 1.12

require (
	github.com/kimpettersen/svc-payments v0.0.0-20190228202415-ed4578bd152a
	google.golang.org/grpc v1.19.0
)

replace github.com/kimpettersen/svc-payments => ../../svc-payments
