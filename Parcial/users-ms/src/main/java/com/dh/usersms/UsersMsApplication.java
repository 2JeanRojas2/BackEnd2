package com.dh.usersms;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.netflix.eureka.EnableEurekaClient;

@SpringBootApplication
@EnableEurekaClient
public class UsersMsApplication {

	public static void main(String[] args) {
		SpringApplication.run(UsersMsApplication.class, args);
	}

}
