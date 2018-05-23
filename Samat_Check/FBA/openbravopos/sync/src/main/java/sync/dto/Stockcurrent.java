package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "stockcurrent")
public class Stockcurrent implements DataModel{

  @Id
  @Column(name = "stockdiary_id")
  private String stockdiary_id;

  @Column(name = "location")
  private String location;
  @Column(name = "product")
  private String product;
  @Column(name = "attributesetinstance_Id")
  private String attributesetinstance_Id;
  @Column(name = "units")
  private double units;
  @Column(name = "sync")
  private long sync;

  public String getStockdiary_id() { return stockdiary_id;}
  public void setStockdiary_id(String stockdiary_id) {
    this.stockdiary_id = stockdiary_id;
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


  public long getSync() {
    return sync;
  }

  public void setSync(long sync) {
    this.sync = sync;
  }

  public String getIdentifier() {return this.stockdiary_id;}

  public String getIdentifierName(){return "stockdiary_id";}

}
