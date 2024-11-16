package br.com.destinara.controller;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;

import br.com.destinara.model.AppUserModel;
import br.com.destinara.repository.AppUserRepository;

@Controller
public class AppUserController {

    private final AppUserRepository appUserRepository;

    public AppUserController(AppUserRepository appUserRepository) {
        this.appUserRepository = appUserRepository;
    }

    @GetMapping("/register")
    public String showRegisterPage() {
        return "register";
    }

    @PostMapping("/register")
    public String registerUser(AppUserModel user) {
        appUserRepository.save(user);
        return "register";
    }

    @GetMapping("/login")
    public String showLoginPage() {
        return "login";
    }

    @PostMapping("/login")
    public String loginUser(String email, String password, Model model) {
        AppUserModel user = appUserRepository.findByEmailAndPassword(email, password).orElse(null);

        if (user != null) {
            return "redirect:/filter?type=Nacional";
        } else {
            model.addAttribute("error", "Email ou senha inválidos");
            return "login";
        }
    }
}
