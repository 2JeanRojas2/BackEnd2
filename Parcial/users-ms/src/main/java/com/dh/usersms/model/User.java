package com.dh.usersms.model;

public class User {
    private String id;
    private String username;
    private String email;
    private String firstName;

    private String nationality;


    public User() {
    }

    public User(String id, String username, String email, String firstName, String nationality) {
        this.id = id;
        this.username = username;
        this.email = email;
        this.firstName = firstName;
        this.nationality = nationality;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getFirstName() {
        return firstName;
    }

    public void setFirstName(String firstName) {
        this.firstName = firstName;
    }

    public String getNationality() {
        return nationality;
    }

    public void setNationality(String nationality) {
        this.nationality = nationality;
    }
}
