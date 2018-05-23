package sync;

import com.google.gson.GsonBuilder;
import com.google.gson.reflect.TypeToken;
import org.hibernate.SQLQuery;
import org.hibernate.Session;
import org.hibernate.Query;

import java.io.*;
import java.lang.reflect.Type;
import java.text.DateFormat;
import java.text.SimpleDateFormat;
import java.util.Calendar;
import java.util.Date;
import java.util.List;

import org.hibernate.SessionFactory;

import com.google.gson.Gson;

import sync.dto.*;

public class App {
    public static Gson gson = null;
    public static Session session = null;
    public static Settings settings = null;
    public static Http http;
    public static String currentDate = getCurrentDate();
    public static String lastDate = getLastDateSync();
    private static String[] tables;
    private static String[] tablesGet;

    public static void main(String[] args) throws IOException {
        //==============================================================================================================
        //                      OPEN DATA BASE CONNECTION
        //==============================================================================================================
        session = HibernateUtil.getSessionFactory().openSession();
        System.out.println("Session statistics" + session.getStatistics());

        //==============================================================================================================
        //                      LOAD CONFIG
        //==============================================================================================================

        gson = new Gson();
        //String file = SessionFactory.class.getClassLoader().getResource("settings.json").getFile();
        //settings = gson.fromJson(new FileReader(file), Settings.class);
        settings = gson.fromJson(new FileReader("settings.json"), Settings.class);
        System.out.println("loading config file" + settings.getKey());

        //==============================================================================================================
        //                      SET DATE
        //==============================================================================================================

        if(lastDate == null) {
            lastDate = "";
        }
        System.out.println("Текущая дата " + currentDate + "Последняя дата получения " + lastDate);

        //==============================================================================================================
        //                      HTTP INITIALIZING
        //==============================================================================================================

        http = new Http(settings.getKey());

        tables  = LoadTables("settables");
        tablesGet = LoadTables("gettables");

        System.out.println("tables = " + tables);
        System.out.println("tablesGet = " + tablesGet);

        //==============================================================================================================
        //                      SEND DATA
        //==============================================================================================================
        for (int i=0; i<tables.length; i++)
        {
            System.out.println("table name = " + tables[i]);
            //get(tables[i]);
        }

        if(isGetTable("Stockcurrent")) {
            System.out.println("Stockcurrents send");
            LoadStockcurrent();
        }
        if(isGetTable("Categories")) {
            LoadCategories();
        }
        if(isGetTable("Taxcategories")) {
            LoadTaxcategories();
        }
        if(isGetTable("Products_Cat")) {
            LoadProducts_Cat();
        }
        if(isGetTable("Products_Com")) {
            LoadProducts_Com();
        }
        if(isGetTable("Products")) {
            LoadProducts();
        }
        //==============================================================================================================
        //                      UPDATE DATA
        //==============================================================================================================

        setLastDateSync(currentDate);


        //==============================================================================================================
        //                      SHUTDOWN
        //==============================================================================================================
        HibernateUtil.shutdown();
    }

    private static String getCurrentDate() {
        DateFormat df = new SimpleDateFormat("MM-dd-yyyy HH:mm:ss:ms");
        Date today = Calendar.getInstance().getTime();
        String reportDate = df.format(today);
        return reportDate.replace(" ", "T");
    }

    private static String getLastDateSync() {
        try {
            File file = new File("lastdate.txt");
            if (!file.exists()) {
                file.createNewFile();
            }
            FileReader fr = new FileReader(file.getAbsoluteFile());
            BufferedReader br = new BufferedReader(fr);
            String date = br.readLine();
            br.close();

            return date;
        } catch (Exception e) {
            e.printStackTrace();
        }
        return "";
    }

    private static void setLastDateSync(String lastDate) {
        try {
            File file = new File("lastdate.txt");
            if (!file.exists()) {
                file.createNewFile();
            }
            FileWriter fw = new FileWriter(file.getAbsoluteFile());
            BufferedWriter bw = new BufferedWriter(fw);
            bw.write(lastDate);
            bw.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    private static void get(String tableName) throws IOException {
        List results = session.createQuery("FROM " + tableName + " WHERE sync = 0").list();
        System.out.println("table " + tableName + "count = " + results.size());
        for (Object a : results) {
            String jsonApplication = new GsonBuilder()
                    .setDateFormat("yyyy-MM-dd'T'hh:mm:ss'Z'")
                    .create()
                    .toJson(a);
            System.out.println("json = " + jsonApplication);
            int code = http.Post(settings.getRestUrl()+tableName.toLowerCase(), jsonApplication);
            DataModel m = (DataModel) a;
            if (code == 201) {
                setSync(tableName, m.getIdentifierName(), m.getIdentifier());
            } else {
                System.out.println("code = " + code);
            }
        }
        getDelete(tableName);
    }

    private static void getDelete(String tableName) throws IOException {
        List<Syncdeletedrecords> deletes = session.createQuery("from Syncdeletedrecords where tablename = '" + tableName.toLowerCase() + "'").list();
        if ( deletes.size() < 1) return;
        for (Syncdeletedrecords a: deletes) {

            int code = http.Delete(settings.getRestUrl()+tableName.toLowerCase() + "/" + a.getPk());
            if (code == 201) {
                deleteRecord(a.getId());
            } else {
                System.out.println("code = " + code);
            }
        }
    }

    public static void setSync(String tableName, String idName, String id){
        session.beginTransaction();
        Query query = session.createQuery("update " + tableName + " set sync = 99" + " where " + idName + " = '" + id +"'");
        query.executeUpdate();
        session.flush();
        session.getTransaction().commit();
    }

    public static void deleteRecord(String id) {
        System.out.println("delet record id = " + id);
        session.beginTransaction();
        Query query = session.createQuery("delete Syncdeletedrecords where id = '" + id + "'");
        query.executeUpdate();
        session.flush();
        session.getTransaction().commit();
    }

    public static void LoadStockcurrent() throws IOException {
        //==============================================================================================================
        //                      GET STOCKDAIRES
        //==============================================================================================================
        String jsonString = http.Get (settings.getRestUrl() + "stockdeary/" + lastDate);
        Type itemsListType = new TypeToken<List<Stockdiary>>(){}.getType();
        List<Stockdiary> stockdiaries = gson.fromJson(jsonString, itemsListType);
        System.out.println("LIST SIZE = " + stockdiaries.size());

        for (Stockdiary s : stockdiaries) {
            System.out.println("INSERT: " + s.toString());

            Stockdiary sd = (Stockdiary)session.load(Stockdiary.class, s.getId());
            System.out.println("SEARCH:" + sd.getId());
            if(sd == null) {
                session.beginTransaction();
                session.save(s);
                session.getTransaction().commit();
            }

        }
        if (stockdiaries.size()>0) {
            session.beginTransaction();
            Query deleteStockcurrent = session.createQuery("delete from Stockcurrent");
            deleteStockcurrent.executeUpdate();
            session.flush();
            SQLQuery insertStockcurrent = session.createSQLQuery("insert into Stockcurrent (stockdiary_id, location, product, units) select (select id from Stockdiary where product=a.product order by datenew desc limit 1), location, product, SUM(units) from Stockdiary a group by product, location");
            insertStockcurrent.executeUpdate();
            session.flush();
            session.getTransaction().commit();
        }
    }

    public static void LoadCategories() throws IOException {
        //==============================================================================================================
        //                      GET CATEGORIES
        //==============================================================================================================
        String jsonString = http.Get (settings.getRestUrl() + "/categories");
        Type itemsListType = new TypeToken<List<Categories>>(){}.getType();

        System.out.println(jsonString);

        List<Categories> categories = gson.fromJson(jsonString, itemsListType);
        //System.out.println("LIST SIZE = " + stockdiaries.size());

        for (Categories c : categories) {


            Categories cd = (Categories)session.load(Categories.class, c.getId());
            //System.out.println("SEARCH:" + sd.getId());
            if(cd == null) {
                System.out.println("INSERT: " + c.toString());
                session.beginTransaction();
                session.save(c);
                session.getTransaction().commit();
            }

        }
    }

    public static void LoadTaxcategories() throws IOException {
        //==============================================================================================================
        //                      GET CATEGORIES
        //==============================================================================================================
        String jsonString = http.Get (settings.getRestUrl() + "/taxcategories");
        Type itemsListType = new TypeToken<List<Taxcategories>>(){}.getType();
        //System.out.println(jsonString);
        List<Taxcategories> taxcategories = gson.fromJson(jsonString, itemsListType);
        //System.out.println("LIST SIZE = " + stockdiaries.size());

        for (Taxcategories t : taxcategories) {
            System.out.println("INSERT: " + t.toString());

            Taxcategories td = (Taxcategories)session.load(Taxcategories.class, t.getId());
            //System.out.println("SEARCH:" + sd.getId());
            if(td == null) {
                session.beginTransaction();
                session.save(t);
                session.getTransaction().commit();
            }

        }
    }

    public static void LoadProducts_Cat() throws IOException {
        //==============================================================================================================
        //                      GET CATEGORIES
        //==============================================================================================================
        String jsonString = http.Get (settings.getRestUrl() + "/products_cat");
        Type products_Cat = new TypeToken<List<Products_Cat>>(){}.getType();
        List<Products_Cat> products_Cats = gson.fromJson(jsonString, products_Cat);
        //System.out.println("LIST SIZE = " + stockdiaries.size());

        for (Products_Cat t : products_Cats) {
            //System.out.println("INSERT: " + s.toString());

            Products_Cat td = (Products_Cat)session.load(Products_Cat.class, t.getProduct());
            //System.out.println("SEARCH:" + sd.getId());
            if(td == null) {
                session.beginTransaction();
                session.save(t);
                session.getTransaction().commit();
            }

        }
    }

    public static void LoadProducts_Com() throws IOException {
        //==============================================================================================================
        //                      GET CATEGORIES
        //==============================================================================================================
        String jsonString = http.Get (settings.getRestUrl() + "/products_com");
        System.out.println("JSONSTRING "+jsonString);
        if(jsonString == "[]") return;
        Type products_Com = new TypeToken<List<Products_Com>>(){}.getType();
        List<Products_Com> products_Coms = gson.fromJson(jsonString, products_Com);
        //System.out.println("LIST SIZE = " + stockdiaries.size());

        for (Products_Com t : products_Coms) {
            //System.out.println("INSERT: " + s.toString());

            Products_Com td = (Products_Com)session.load(Products_Com.class, t.getProduct());
            //System.out.println("SEARCH:" + sd.getId());
            if(td == null) {
                session.beginTransaction();
                session.save(t);
                session.getTransaction().commit();
            }

        }
    }

    public static void LoadProducts() throws IOException {
        //==============================================================================================================
        //                      GET CATEGORIES
        //==============================================================================================================
        String jsonString = http.Get (settings.getRestUrl() + "products");
        System.out.println(jsonString);
        Type products = new TypeToken<List<Products>>(){}.getType();
        List<Products> productsList = gson.fromJson(jsonString, products);
        //System.out.println("LIST SIZE = " + stockdiaries.size());

        for (Products t : productsList) {
            //System.out.println("INSERT: " + s.toString());

            Products td = (Products)session.load(Products.class, t.getId());
            //System.out.println("SEARCH:" + sd.getId());
            if(td == null) {
                session.beginTransaction();
                session.save(t);
                session.getTransaction().commit();
            }

        }
    }

    public static String[] LoadTables(String command) throws IOException {
        //==============================================================================================================
        //                      GET CATEGORIES
        //==============================================================================================================
        String jsonString = http.Get (settings.getRestUrl() + command);
        System.out.println(jsonString);

        jsonString = jsonString.replace("[","");
        jsonString = jsonString.replace("]","");
        return jsonString.split(" ");


    }

    public static boolean isGetTable(String tableName) {
        for (String a : tablesGet) {
            if(a.equals(tableName)) return true;
        }
        return false;
    }

}
