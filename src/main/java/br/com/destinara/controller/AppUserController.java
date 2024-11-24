package br.com.destinara.controller;

import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.User;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import br.com.destinara.model.AppUserModel;
import br.com.destinara.repository.AppUserRepository;

@Controller
public class AppUserController {

    private final AppUserRepository appUserRepository;
    private final PasswordEncoder passwordEncoder;

    public AppUserController(AppUserRepository appUserRepository, PasswordEncoder passwordEncoder) {
        this.appUserRepository = appUserRepository;
        this.passwordEncoder = passwordEncoder;
    }

    @GetMapping("/register")
    public String showRegisterPage() {
        return "register";
    }

    @PostMapping("/register")
    public String registerUser(AppUserModel user) {
        user.setPassword(passwordEncoder.encode(user.getPassword()));
        appUserRepository.save(user);
        return "register";
    }

    @GetMapping("/login")
    public String showLoginPage() {
        return "login";
    }

    @PostMapping("/login")
    public String loginUser(String email, String password, Model model) {
    return "login";
    }

    @GetMapping("/edit-user")
    public String showEditUserPage(Model model) {

    UserDetails userDetails = (UserDetails) SecurityContextHolder.getContext().getAuthentication().getPrincipal();
    AppUserModel user = appUserRepository.findByEmail(userDetails.getUsername()).orElse(null);

    if (user != null) {
        model.addAttribute("user", user);
        return "edit-user";
    }
        return "redirect:/purchase-history";
    }

     @PostMapping("/update-user")
    public String updateUser(AppUserModel updatedUser) {
        UserDetails userDetails = (UserDetails) SecurityContextHolder.getContext().getAuthentication().getPrincipal();
        AppUserModel currentUser = appUserRepository.findByEmail(userDetails.getUsername()).orElse(null);

        if (currentUser != null) {
            currentUser.setName(updatedUser.getName());
            currentUser.setEmail(updatedUser.getEmail());
            currentUser.setCpf(updatedUser.getCpf());
            currentUser.setBirthDate(updatedUser.getBirthDate());
            currentUser.setSex(updatedUser.getSex());

            appUserRepository.save(currentUser);

            User updatedUserDetails = new User(
                currentUser.getEmail(), 
                currentUser.getPassword(),
                userDetails.getAuthorities()
            );
            UsernamePasswordAuthenticationToken newAuth = new UsernamePasswordAuthenticationToken(
                updatedUserDetails, 
                null, 
                updatedUserDetails.getAuthorities()
            );
            SecurityContextHolder.getContext().setAuthentication(newAuth);
        }
        return "redirect:/purchase-history"; 
    }
}