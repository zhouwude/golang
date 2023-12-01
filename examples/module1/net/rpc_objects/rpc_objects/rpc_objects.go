// rpc_objects.go
package rpc_objects

type Args struct {
	N, M int
}

//reply *int 通过参数改变值 传递的是指针而不是 int 是拷贝无法修改原来的值
func (t *Args) Multiply(args *Args, reply *int) error {
	*reply = args.N * args.M
	return nil
}
