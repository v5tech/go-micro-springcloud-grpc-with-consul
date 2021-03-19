package com.anchnet.grpc.grpc.server;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;
import org.springframework.web.bind.annotation.RestController;

@RestController
@EnableDiscoveryClient
@SpringBootApplication
public class GrpcServerWithConsulApplication {

	public static void main(String[] args) {
		SpringApplication.run(GrpcServerWithConsulApplication.class, args);
	}

}
