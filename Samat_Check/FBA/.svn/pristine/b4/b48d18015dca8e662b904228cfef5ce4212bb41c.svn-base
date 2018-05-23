package sync.dto;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;
import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "stockdiary", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Stockdiary implements DataModel{
  @Override
  public String toString() {
    return "Stockdiary{" +
            "id='" + id + '\'' +
            ", datenew=" + datenew +
            ", reason=" + reason +
            ", location='" + location + '\'' +
            ", product='" + product + '\'' +
            ", attributesetinstance_Id='" + attributesetinstance_Id + '\'' +
            ", units=" + units +
            ", price=" + price +
            ", sync=" + sync +
            '}';
  }

  @Id
  @Column(name = "id", unique = true, nullable = false)
  @SerializedName("Id")
  @Expose
  private String id;

  @SerializedName("Datenew")
  @Expose
  @Column(name = "datenew")
  private java.sql.Timestamp datenew;

  @SerializedName("Reason")
  @Expose
  @Column(name = "reason")
  private long reason;

  @SerializedName("Location")
  @Expose
  @Column(name = "location")
  private String location;

  @SerializedName("Product")
  @Expose
  @Column(name = "product")
  private String product;

  @SerializedName("Attributesetinstance_Id")
  @Expose
  @Column(name = "attributesetinstance_Id")
  private String attributesetinstance_Id;

  @SerializedName("Units")
  @Expose
  @Column(name = "units")
  private double units;

  @SerializedName("Price")
  @Expose
  @Column(name = "price")
  private double price;

  @Column(name = "sync")
  private long sync;


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


  public long getReason() {
    return reason;
  }

  public void setReason(long reason) {
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


  public String getAttributesetinstance_Id() {
    return attributesetinstance_Id;
  }

  public void setAttributesetinstance_Id(String attributesetinstance_Id) {
    this.attributesetinstance_Id = attributesetinstance_Id;
  }


  public double getUnits() {
    return units;
  }

  public void setUnits(double units) {
    this.units = units;
  }


  public double getPrice() {
    return price;
  }

  public void setPrice(double price) {
    this.price = price;
  }


  public long getSync() {
    return sync;
  }

  public void setSync(long sync) {
    this.sync = sync;
  }

  public String getIdentifier() {return this.id;}

  public String getIdentifierName(){return "id";}

}
