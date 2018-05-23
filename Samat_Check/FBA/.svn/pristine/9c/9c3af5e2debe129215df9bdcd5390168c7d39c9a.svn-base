package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "customers", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Customers implements DataModel{

  @Id
  @Column(name = "id", unique = true, nullable = false)
  private String id;
  @Column(name = "searchkey")
  private String searchkey;
  @Column(name = "taxid")
  private String taxid;
  @Column(name = "name")
  private String name;
  @Column(name = "taxcategory")
  private String taxcategory;
  @Column(name = "card")
  private String card;
  @Column(name = "maxdebt")
  private double maxdebt;
  @Column(name = "address")
  private String address;
  @Column(name = "address2")
  private String address2;
  @Column(name = "postal")
  private String postal;
  @Column(name = "city")
  private String city;
  @Column(name = "region")
  private String region;
  @Column(name = "country")
  private String country;
  @Column(name = "firstname")
  private String firstname;
  @Column(name = "lastname")
  private String lastname;
  @Column(name = "email")
  private String email;
  @Column(name = "phone")
  private String phone;
  @Column(name = "phone2")
  private String phone2;
  @Column(name = "fax")
  private String fax;
  @Column(name = "notes")
  private String notes;
  @Column(name = "visible")
  private String visible;
  @Column(name = "curdate")
  private java.sql.Timestamp curdate;
  @Column(name = "curdebt")
  private float curdebt;
  @Column(name = "sync")
  private long sync;


  public String getId() {
    return id;
  }

  public void setId(String id) {
    this.id = id;
  }


  public String getSearchkey() {
    return searchkey;
  }

  public void setSearchkey(String searchkey) {
    this.searchkey = searchkey;
  }


  public String getTaxid() {
    return taxid;
  }

  public void setTaxid(String taxid) {
    this.taxid = taxid;
  }


  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }


  public String getTaxcategory() {
    return taxcategory;
  }

  public void setTaxcategory(String taxcategory) {
    this.taxcategory = taxcategory;
  }


  public String getCard() {
    return card;
  }

  public void setCard(String card) {
    this.card = card;
  }


  public double getMaxdebt() {
    return maxdebt;
  }

  public void setMaxdebt(double maxdebt) {
    this.maxdebt = maxdebt;
  }


  public String getAddress() {
    return address;
  }

  public void setAddress(String address) {
    this.address = address;
  }


  public String getAddress2() {
    return address2;
  }

  public void setAddress2(String address2) {
    this.address2 = address2;
  }


  public String getPostal() {
    return postal;
  }

  public void setPostal(String postal) {
    this.postal = postal;
  }


  public String getCity() {
    return city;
  }

  public void setCity(String city) {
    this.city = city;
  }


  public String getRegion() {
    return region;
  }

  public void setRegion(String region) {
    this.region = region;
  }


  public String getCountry() {
    return country;
  }

  public void setCountry(String country) {
    this.country = country;
  }


  public String getFirstname() {
    return firstname;
  }

  public void setFirstname(String firstname) {
    this.firstname = firstname;
  }


  public String getLastname() {
    return lastname;
  }

  public void setLastname(String lastname) {
    this.lastname = lastname;
  }


  public String getEmail() {
    return email;
  }

  public void setEmail(String email) {
    this.email = email;
  }


  public String getPhone() {
    return phone;
  }

  public void setPhone(String phone) {
    this.phone = phone;
  }


  public String getPhone2() {
    return phone2;
  }

  public void setPhone2(String phone2) {
    this.phone2 = phone2;
  }


  public String getFax() {
    return fax;
  }

  public void setFax(String fax) {
    this.fax = fax;
  }


  public String getNotes() {
    return notes;
  }

  public void setNotes(String notes) {
    this.notes = notes;
  }


  public String getVisible() {
    return visible;
  }

  public void setVisible(String visible) {
    this.visible = visible;
  }


  public java.sql.Timestamp getCurdate() {
    return curdate;
  }

  public void setCurdate(java.sql.Timestamp curdate) {
    this.curdate = curdate;
  }


  public float getCurdebt() {
    return curdebt;
  }

  public void setCurdebt(float curdebt) {
    //if(null != Double.valueOf(curdebt)){
    if(Float.valueOf(curdebt) != null){
      this.curdebt = curdebt;
    }else{
      this.curdebt = 0;
    }
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
