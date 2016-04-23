default:
	go build dsproject/carnode
	go build dsproject/server
	go build dsproject/supernode

clean:
	rm carnode server supernode
