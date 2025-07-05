package br.com.destinara.repository;

import br.com.destinara.model.AppUserModel;
import br.com.destinara.model.PurchaseModel;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.List;

public interface PurchaseRepository extends JpaRepository<PurchaseModel, Integer> {
    List<PurchaseModel> findByUser(AppUserModel user);

    @Query("SELECT p.travelPackage.name, COUNT(p.travelPackage) FROM PurchaseModel p GROUP BY p.travelPackage ORDER BY COUNT(p.travelPackage) DESC")
    List<Object[]> findTopSellingPackages();
}
