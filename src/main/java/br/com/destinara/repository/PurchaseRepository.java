package br.com.destinara.repository;

import br.com.destinara.model.PurchaseModel;
import org.springframework.data.jpa.repository.JpaRepository;

public interface PurchaseRepository extends JpaRepository<PurchaseModel, Integer> {
}
