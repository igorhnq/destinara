package br.com.destinara.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import br.com.destinara.model.TravelPackageModel;

public interface TravelPackageRepository extends JpaRepository <TravelPackageModel, Integer> {
        
}