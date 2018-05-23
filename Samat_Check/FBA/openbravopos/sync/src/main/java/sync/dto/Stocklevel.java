package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "stocklevel", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Stocklevel implements DataModel{

  @Id
  @Column(name = "id", unique = true, nullable = false)
  private String id;
  @Column(name = "location")
  private String location;
  @Column(name = "product")
  private String product;
  @Column(name = "stocksecurity")
  private double stocksecurity;
  @Column(name = "stockmaximum")
  private double stockmaximum;
  @Column(name = "sync")
  private long sync;


  public String getId() {
    return id;
  }

  public void setId(String id) {
    this.id = id;
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


  public double getStocksecurity() {
    return stocksecurity;
  }

  public void setStocksecurity(double stocksecurity) {
    this.stocksecurity = stocksecurity;
  }


  public double getStockmaximum() {
    return stockmaximum;
  }

  public void setStockmaximum(double stockmaximum) {
    this.stockmaximum = stockmaximum;
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
