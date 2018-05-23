package sync.dto;

import com.google.gson.annotations.SerializedName;
import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "products_Com", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Products_Com implements DataModel{

  @Id
  @Column(name = "id", unique = true, nullable = false)
  @SerializedName("Id")
  private String id;
  @Column(name = "product")
  @SerializedName("Product")
  private String product;
  @Column(name = "product2")
  @SerializedName("Product2")
  private String product2;
  @Column(name = "sync")
  private long sync;


  public String getId() {
    return id;
  }

  public void setId(String id) {
    this.id = id;
  }


  public String getProduct() {
    return product;
  }

  public void setProduct(String product) {
    this.product = product;
  }


  public String getProduct2() {
    return product2;
  }

  public void setProduct2(String product2) {
    this.product2 = product2;
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
