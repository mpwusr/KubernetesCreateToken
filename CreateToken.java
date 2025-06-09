import io.fabric8.kubernetes.api.model.TokenRequest;
import io.fabric8.kubernetes.api.model.TokenRequestSpec;
import io.fabric8.kubernetes.api.model.authentication.TokenRequestBuilder;
import io.fabric8.kubernetes.client.DefaultKubernetesClient;
import io.fabric8.kubernetes.client.KubernetesClient;

public class GenerateSAToken {
    public static void main(String[] args) {
        try (KubernetesClient client = new DefaultKubernetesClient()) {
            TokenRequest tokenRequest = new TokenRequestBuilder()
                .withNewMetadata().endMetadata()
                .withSpec(new TokenRequestSpec(null, 3600L, java.util.List.of("https://kubernetes.default.svc"), null))
                .build();

            TokenRequest token = client.serviceAccounts()
                .inNamespace("default")
                .withName("podcast")
                .createToken(tokenRequest);

            System.out.println("Token: " + token.getStatus().getToken());
        }
    }
}
