package br.com.destinara.controller;

import br.com.destinara.model.AppUserModel;
import br.com.destinara.model.PurchaseModel;
import br.com.destinara.model.TravelPackageModel;
import br.com.destinara.repository.AppUserRepository;
import br.com.destinara.repository.PurchaseRepository;
import br.com.destinara.repository.TravelPackageRepository;

import java.util.List;

import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;

@Controller
public class TravelPackageController {

    private final TravelPackageRepository travelPackageRepository;
    private final AppUserRepository appUserRepository;
    private final PurchaseRepository purchaseRepository;

    public TravelPackageController(TravelPackageRepository travelPackageRepository, 
                                   AppUserRepository appUserRepository, 
                                   PurchaseRepository purchaseRepository) {
        this.travelPackageRepository = travelPackageRepository;
        this.appUserRepository = appUserRepository;
        this.purchaseRepository = purchaseRepository;
    }

    @GetMapping("/purchase-history")
    public String viewPurchaseHistory(Model model) {
        UserDetails userDetails = (UserDetails) SecurityContextHolder.getContext().getAuthentication().getPrincipal();
        AppUserModel user = appUserRepository.findByEmail(userDetails.getUsername()).orElse(null);
    
        if (user != null) {
            model.addAttribute("user", user);
    
            List<PurchaseModel> purchases = purchaseRepository.findByUser(user);
            model.addAttribute("purchases", purchases);
        }
    
        return "purchase-history"; 
    }
    
    @GetMapping("/travel-packages")
    public String showPackages(Model model) {
        List<TravelPackageModel> travelPackageModels = travelPackageRepository.findAll();
        model.addAttribute("travelPackages", travelPackageModels);
        return "index";
    }

    @GetMapping("/filter")
    public String filterByType(@RequestParam String type, Model model) {
        List<TravelPackageModel> filteredPackages = travelPackageRepository.findByType(type);
        model.addAttribute("travelPackages", filteredPackages);
        return "packages";
    }

    @GetMapping("/package-details")
    public String showPackageDetails(@RequestParam Integer id, Model model) {
        TravelPackageModel travelPackage = travelPackageRepository.findById(id).orElse(null);
        model.addAttribute("travelPackage", travelPackage);
        return "package-details";
    }

    @GetMapping("/buy-package")
    public String buyPackage(@RequestParam Integer packageId, Model model) {
        UserDetails userDetails = (UserDetails) SecurityContextHolder.getContext().getAuthentication().getPrincipal();
        AppUserModel user = appUserRepository.findByEmail(userDetails.getUsername()).orElse(null);

        if (user != null) {
            TravelPackageModel travelPackage = travelPackageRepository.findById(packageId).orElse(null);
            if (travelPackage != null) {
                PurchaseModel purchase = new PurchaseModel();
                purchase.setUser(user);
                purchase.setTravelPackage(travelPackage);
                purchaseRepository.save(purchase);
                
                model.addAttribute("message", "Compra realizada com sucesso!");
            }
        }

        return "purchase-success";
    }
}
