export UDIR= .
export GOC = x86_64-xen-ethos-6g
export GOL = x86_64-xen-ethos-6l
export ETN2GO = etn2go
export ET2G   = et2g
export EG2GO  = eg2go

export GOARCH = amd64
export TARGET_ARCH = x86_64
export GOETHOSINCLUDE=/usr/lib64/go/pkg/ethos_$(GOARCH)
export GOLINUXINCLUDE=/usr/lib64/go/pkg/linux_$(GOARCH)

export ETHOSROOT=server/rootfs
export MINIMALTDROOT=server/minimaltdfs

.PHONY: all install
all: serverbasic clientbasic

myRpc.go: myRpc.t
	$(ETN2GO) . myRpc main $^

serverbasic: serverbasic.go myRpc.go
	ethosGo $^

clientbasic: clientbasic.go myRpc.go
	ethosGo $^

install: clean serverbasic clientbasic
	(ethosParams server && cd server && ethosMinimaltdBuilder)
	ethosTypeInstall myRpc
	ethosServiceInstall myRpc
	ethosDirCreate $(ETHOSROOT)/services/myRpc   $(ETHOSROOT)/types/spec/myRpc/MyRpc all
	install -D  serverbasic clientbasic	$(ETHOSROOT)/programs
	ethosStringEncode /programs/serverbasic > $(ETHOSROOT)/etc/init/services/serverbasic
	ethosStringEncode /programs/clientbasic > $(ETHOSROOT)/etc/init/services/clientbasic

clean:
	sudo rm -rf server
	rm -rf myRpc/ myRpcIndex/
	rm -f myRpc.go
	rm -f serverbasic
	rm -f serverbasic.goo.ethos
	rm -f clientbasic
	rm -f clientbasic.goo.ethos
