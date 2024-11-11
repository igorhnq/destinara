package br.com.destinara.controller;

import java.util.List;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

import br.com.destinara.model.TravelPackageModel;
import br.com.destinara.repository.TravelPackageRepository;

@Controller
public class TravelPackageController {
    
    private final TravelPackageRepository travelPackageRepository;

    public TravelPackageController(TravelPackageRepository travelPackageRepository) {
        this.travelPackageRepository = travelPackageRepository;
    }

    @GetMapping("/travel-packages")
    public String showPackages(Model model) {
        List<TravelPackageModel> travelPackageModels = travelPackageRepository.findAll();
        model.addAttribute("travelPackages", travelPackageModels);
        return "index";
    }
}
