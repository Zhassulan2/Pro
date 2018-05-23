package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "taxes", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Taxes implements DataModel{

  @Id
  @Column(name = "id", unique = true, nullable = false)
  private String id;
  @Column(name = "name")
  private String name;
  @Column(name = "category")
  private String category;
  @Column(name = "custcategory")
  private String custcategory;
  @Column(name = "parentid")
  private String parentid;
  @Column(name = "rate")
  private double rate;
  @Column(name = "ratecascade")
  private String ratecascade;
  @Column(name = "rateorder")
  private java.lang.Long rateorder;
  //@Column(name = "sync")
  //private transient long sync;


  public String getId() {
    return id;
  }

  public void setId(String id) {
    this.id = id;
  }


  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }


  public String getCategory() {
    return category;
  }

  public void setCategory(String category) {
    this.category = category;
  }


  public String getCustcategory() {
    return custcategory;
  }

  public void setCustcategory(String custcategory) {
    this.custcategory = custcategory;
  }


  public String getParentid() {
    return parentid;
  }

  public void setParentid(String parentid) {
    this.parentid = parentid;
  }


  public double getRate() {
    return rate;
  }

  public void setRate(double rate) {
    this.rate = rate;
  }


  public String getRatecascade() {
    return ratecascade;
  }

  public void setRatecascade(String ratecascade) {
    this.ratecascade = ratecascade;
  }


  public java.lang.Long getRateorder() {
    return rateorder;
  }

  public void setRateorder(java.lang.Long rateorder) {
    this.rateorder = rateorder;
  }

/*
  public long getSync() {
    return sync;
  }

  public void setSync(long sync) {
    this.sync = sync;
  }
*/
  public String getIdentifier() {return this.id;}

  public String getIdentifierName(){return "id";}

}
