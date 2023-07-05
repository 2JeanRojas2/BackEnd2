package com.dh.usersms.controller;

import com.dh.usersms.model.User;
import com.dh.usersms.service.UserService;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/users")
public class UserRestController {
    private final UserService userService;

    public UserRestController(UserService userService) {
        this.userService = userService;
    }

    @GetMapping("/firstname/{firstName}")
    public List<User> findByFirstName(@PathVariable String firstName) {
        return userService.findByFirstName(firstName);
    }

    @GetMapping("/id/{id}")
    public User findById(@PathVariable String id) {
        return userService.findById(id);
    }

    @PutMapping("/update")
    public User findById(@RequestParam String id, @RequestParam String nationality) {
        return userService.updateNationality(id, nationality);
    }

}
