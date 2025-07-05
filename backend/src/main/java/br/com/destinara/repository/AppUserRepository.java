package br.com.destinara.repository;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;

import br.com.destinara.model.AppUserModel;

public interface AppUserRepository extends JpaRepository<AppUserModel, Integer> {
    Optional<AppUserModel> findByEmail(String email);
}
