package com.dh.usersms.repository;

import com.dh.usersms.model.User;

import java.util.List;

public interface IUserRepository {

    List<User> findByFirstName(String name);

    User findById(String id);

    User updateNationality(String id, String nationality);
}
