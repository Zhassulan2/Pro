package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "locations", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Locations implements DataModel{

  @Id
  @Column(name = "id", unique = true, nullable = false)
  private String id;
  @Column(name = "name")
  private String name;
  @Column(name = "address")
  private String address;
  @Column(name = "sync")
  private long sync;


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


  public String getAddress() {
    return address;
  }

  public void setAddress(String address) {
    this.address = address;
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
