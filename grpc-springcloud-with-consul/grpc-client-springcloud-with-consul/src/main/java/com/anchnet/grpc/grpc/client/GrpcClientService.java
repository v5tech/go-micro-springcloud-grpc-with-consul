package com.anchnet.grpc.grpc.client;

import com.anchnet.grpc.proto.HelloGrpc;
import com.anchnet.grpc.proto.HelloReply;
import com.anchnet.grpc.proto.HelloRequest;
import net.devh.boot.grpc.client.inject.GrpcClient;
import org.springframework.stereotype.Service;

@Service
public class GrpcClientService {

    /**
     * golang_grpc_server为服务名
     */
    @GrpcClient("golang_grpc_server")
    HelloGrpc.HelloBlockingStub golangGrpcServer;

    /**
     * grpc-server-springcloud-with-consul为服务名
     */
    @GrpcClient("grpc-server-springcloud-with-consul")
    HelloGrpc.HelloBlockingStub helloBlockingStub;

    public String sendMessage(String name) {
        try {
            HelloReply helloReply = helloBlockingStub.sayHello(HelloRequest.newBuilder().setName(name).build());
            return helloReply.getMessage();
        } catch (Exception e) {
            e.printStackTrace();
            System.out.println(e.getMessage());
            return "FAILED with " + e.getMessage();
        }
    }

}
