package br.com.destinara.repository;

import br.com.destinara.model.AppUserModel;
import br.com.destinara.model.PurchaseModel;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface PurchaseRepository extends JpaRepository<PurchaseModel, Integer> {
    List<PurchaseModel> findByUser(AppUserModel user);
}
