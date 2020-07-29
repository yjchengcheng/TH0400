package errors

// ERRCODE ...
type ERRCODE int

// error code const define
const (
	OK            ERRCODE = 200
	BadRequestErr ERRCODE = 400
	InternalErr   ERRCODE = 500

	AuthNeedErr ERRCODE = 0400
)

// error code const define
const (
	_ ERRCODE = iota + 10000
)

// ErrMap ...
var ErrMap map[ERRCODE]string = map[ERRCODE]string{
	OK:            "okok no bb",
	BadRequestErr: "request data is not valid",
	InternalErr:   "server internal error",

	AuthNeedErr: "must login or fuck you",
}
