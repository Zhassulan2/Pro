package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "thirdparties", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Thirdparties implements DataModel{
  @Id
  @Column(name = "id", unique = true, nullable = false)
  private String id;
  @Column(name = "cif")
  private String cif;
  @Column(name = "name")
  private String name;
  @Column(name = "address")
  private String address;
  @Column(name = "contactcomm")
  private String contactcomm;
  @Column(name = "contactfact")
  private String contactfact;
  @Column(name = "payrule")
  private String payrule;
  @Column(name = "faxnumber")
  private String faxnumber;
  @Column(name = "phonenumber")
  private String phonenumber;
  @Column(name = "mobilenumber")
  private String mobilenumber;
  @Column(name = "email")
  private String email;
  @Column(name = "webpage")
  private String webpage;
  @Column(name = "notes")
  private String notes;
  @Column(name = "sync")
  private long sync;


  public String getId() {
    return id;
  }

  public void setId(String id) {
    this.id = id;
  }


  public String getCif() {
    return cif;
  }

  public void setCif(String cif) {
    this.cif = cif;
  }


  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }


  public String getAddress() {
    return address;
  }

  public void setAddress(String address) {
    this.address = address;
  }


  public String getContactcomm() {
    return contactcomm;
  }

  public void setContactcomm(String contactcomm) {
    this.contactcomm = contactcomm;
  }


  public String getContactfact() {
    return contactfact;
  }

  public void setContactfact(String contactfact) {
    this.contactfact = contactfact;
  }


  public String getPayrule() {
    return payrule;
  }

  public void setPayrule(String payrule) {
    this.payrule = payrule;
  }


  public String getFaxnumber() {
    return faxnumber;
  }

  public void setFaxnumber(String faxnumber) {
    this.faxnumber = faxnumber;
  }


  public String getPhonenumber() {
    return phonenumber;
  }

  public void setPhonenumber(String phonenumber) {
    this.phonenumber = phonenumber;
  }


  public String getMobilenumber() {
    return mobilenumber;
  }

  public void setMobilenumber(String mobilenumber) {
    this.mobilenumber = mobilenumber;
  }


  public String getEmail() {
    return email;
  }

  public void setEmail(String email) {
    this.email = email;
  }


  public String getWebpage() {
    return webpage;
  }

  public void setWebpage(String webpage) {
    this.webpage = webpage;
  }


  public String getNotes() {
    return notes;
  }

  public void setNotes(String notes) {
    this.notes = notes;
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
