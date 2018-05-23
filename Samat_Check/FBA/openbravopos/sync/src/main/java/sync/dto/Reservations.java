package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "reservations", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Reservations implements DataModel{

  @Id
  @Column(name = "id", unique = true, nullable = false)
  private String id;
  @Column(name = "created")
  private java.sql.Timestamp created;
  @Column(name = "datenew")
  private java.sql.Timestamp datenew;
  @Column(name = "title")
  private String title;
  @Column(name = "chairs")
  private long chairs;
  @Column(name = "isdone")
  private String isdone;
  @Column(name = "description")
  private String description;
  @Column(name = "sync")
  private long sync;


  public String getId() {
    return id;
  }

  public void setId(String id) {
    this.id = id;
  }


  public java.sql.Timestamp getCreated() {
    return created;
  }

  public void setCreated(java.sql.Timestamp created) {
    this.created = created;
  }


  public java.sql.Timestamp getDatenew() {
    return datenew;
  }

  public void setDatenew(java.sql.Timestamp datenew) {
    this.datenew = datenew;
  }


  public String getTitle() {
    return title;
  }

  public void setTitle(String title) {
    this.title = title;
  }


  public long getChairs() {
    return chairs;
  }

  public void setChairs(long chairs) {
    this.chairs = chairs;
  }


  public String getIsdone() {
    return isdone;
  }

  public void setIsdone(String isdone) {
    this.isdone = isdone;
  }


  public String getDescription() {
    return description;
  }

  public void setDescription(String description) {
    this.description = description;
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
