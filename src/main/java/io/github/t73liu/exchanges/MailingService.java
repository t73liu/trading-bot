package io.github.t73liu.exchanges;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.mail.SimpleMailMessage;
import org.springframework.mail.javamail.JavaMailSenderImpl;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;

@Service
@ConfigurationProperties(prefix = "MailingService")
public class MailingService {
    private static final Logger LOGGER = LoggerFactory.getLogger(MailingService.class);

    private final JavaMailSenderImpl mailer;

    private List<String> recipients = new ArrayList<>();

    @Autowired
    private MailingService(JavaMailSenderImpl mailer) {
        this.mailer = mailer;
    }

    public List<String> getRecipients() {
        return recipients;
    }

    public void sendMail(String subject, String text) {
        SimpleMailMessage message = new SimpleMailMessage();
        message.setFrom(mailer.getUsername());
        message.setSubject(subject);
        message.setText(text);
        message.setTo(recipients.toArray(new String[recipients.size()]));
        LOGGER.info("Sending mail: {}, to: {}", message, recipients);
        mailer.send(message);
    }
}
