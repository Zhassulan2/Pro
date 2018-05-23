package sync;

import com.google.gson.reflect.TypeToken;
import org.apache.http.HttpResponse;
import org.apache.http.client.HttpClient;
import org.apache.http.client.methods.HttpDelete;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.entity.StringEntity;
import org.apache.http.impl.client.HttpClientBuilder;
import org.apache.http.protocol.HTTP;
import org.apache.http.util.EntityUtils;
import sync.dto.Stockdiary;

import java.io.IOException;
import java.lang.reflect.Type;
import java.util.Arrays;
import java.util.List;

import com.google.gson.Gson;

public class Http {

    private String key;

    public Http(String key) {
        this.key = key;
    }

    public int Post(String url, String body) throws IOException {
        try {
            System.out.println("start send " + key + url + body);
            HttpClient httpClient    = HttpClientBuilder.create().build();
            HttpPost post          = new HttpPost(url);
            StringEntity postingString = new StringEntity(body, HTTP.UTF_8);
            post.setEntity(postingString);
            post.setHeader("Content-type", "application/json");
            post.setHeader("Accept", "application/json");
            post.setHeader("charset", "utf-8");
            post.setHeader("key", key);

            HttpResponse response = httpClient.execute(post);
            return response.getStatusLine().getStatusCode();
        } catch(Exception ex) {
            System.out.println("ERROR:" + ex.toString());
        }
        return -1;
    }

    public int Delete(String url) throws IOException {
        try {
            System.out.println("start send " + url);
            HttpClient   httpClient    = HttpClientBuilder.create().build();
            HttpDelete delete        = new HttpDelete(url);
            delete.setHeader("Content-type", "application/json");
            delete.setHeader("Accept", "application/json");
            delete.setHeader("charset", "utf-8");
            delete.setHeader("key", key);
            HttpResponse response = httpClient.execute(delete);
            return response.getStatusLine().getStatusCode();
        } catch(Exception ex) {
            System.out.println("ERROR:" + ex.toString());
        }
        return -1;
    }

    public String Get(String url) throws IOException {
        Gson gson = new Gson();

        try {
            System.out.println("start send " + url);
            HttpClient   httpClient    = HttpClientBuilder.create().build();
            HttpGet get        = new HttpGet(url);
            get.setHeader("Content-type", "application/json");
            get.setHeader("Accept", "application/json");
            get.setHeader("charset", "utf-8");
            get.setHeader("key", key);
            HttpResponse response = httpClient.execute(get);
            String json_string = EntityUtils.toString(response.getEntity());
            return json_string;
            /*System.out.println(json_string);


            Type itemsListType = new TypeToken<List<Stockdiary>>(){}.getType();
            List<Stockdiary> stockdiaries = gson.fromJson(json_string, itemsListType);
            System.out.println("LIST SIZE = " + stockdiaries.size());


            for (Stockdiary s : stockdiaries) {
                System.out.println(s.toString());
            }*/

        } catch(Exception ex) {
            System.out.println("ERROR:" + ex.toString());
        }
        return null;
    }
}
