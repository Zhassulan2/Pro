package sync;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class StockdearyJSON {

    @SerializedName("Id")
    @Expose
    private String id;
    @SerializedName("Datenew")
    @Expose
    private java.sql.Timestamp datenew;
    @SerializedName("Reason")
    @Expose
    private Integer reason;
    @SerializedName("Location")
    @Expose
    private String location;
    @SerializedName("Product")
    @Expose
    private String product;
    @SerializedName("Attributesetinstance_Id")
    @Expose
    private String attributesetinstanceId;
    @SerializedName("Units")
    @Expose
    private Integer units;
    @SerializedName("Price")
    @Expose
    private Integer price;

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public java.sql.Timestamp getDatenew() {
        return datenew;
    }

    public void setDatenew(java.sql.Timestamp datenew) {
        this.datenew = datenew;
    }

    public Integer getReason() {
        return reason;
    }

    public void setReason(Integer reason) {
        this.reason = reason;
    }

    public String getLocation() {
        return location;
    }

    public void setLocation(String location) {
        this.location = location;
    }

    public String getProduct() {
        return product;
    }

    public void setProduct(String product) {
        this.product = product;
    }

    public String getAttributesetinstanceId() {
        return attributesetinstanceId;
    }

    public void setAttributesetinstanceId(String attributesetinstanceId) {
        this.attributesetinstanceId = attributesetinstanceId;
    }

    public Integer getUnits() {
        return units;
    }

    public void setUnits(Integer units) {
        this.units = units;
    }

    public Integer getPrice() {
        return price;
    }

    public void setPrice(Integer price) {
        this.price = price;
    }

    @Override
    public String toString() {
        return "StockdearyJSON{" +
                "id='" + id + '\'' +
                ", datenew='" + datenew + '\'' +
                ", reason=" + reason +
                ", location='" + location + '\'' +
                ", product='" + product + '\'' +
                ", attributesetinstanceId=" + attributesetinstanceId +
                ", units=" + units +
                ", price=" + price +
                '}';
    }
}