module Modgo

go 1.14

require github.com/gofiber/fiber/v2 v2.1.4

require (
	Modgo/domain v0.0.0
	github.com/google/uuid v1.1.2
)

replace Modgo/domain v0.0.0 => ../domain
