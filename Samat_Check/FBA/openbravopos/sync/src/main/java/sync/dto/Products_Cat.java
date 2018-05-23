package sync.dto;

import com.google.gson.annotations.SerializedName;
import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "products_Cat")
public class Products_Cat implements DataModel{

  @Id
  @Column(name = "product")
  @SerializedName("Product")
  private String product;
  @Column(name = "catorder")
  @SerializedName("Catorder")
  private java.lang.Long catorder;
  @Column(name = "sync")
  private long sync;


  public String getProduct() {
    return product;
  }

  public void setProduct(String product) {
    this.product = product;
  }


  public java.lang.Long getCatorder() {
    return catorder;
  }

  public void setCatorder(java.lang.Long catorder) {
    this.catorder = catorder;
  }


  public long getSync() {
    return sync;
  }

  public void setSync(long sync) {
    this.sync = sync;
  }

  public String getIdentifier() {return this.product;}

  public String getIdentifierName(){return "product";}

}
