package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@SuppressWarnings("deprecation")
@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "attribute", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Attribute implements DataModel{

  @Id
  @Column(name = "id", unique = true, nullable = false)
  private String id;
  @Column(name = "name")
  private String name;
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


  public long getSync() {
    return sync;
  }

  public void setSync(long sync) {
    this.sync = sync;
  }

  public String getIdentifier() {
    return this.id;
  }
    public String getIdentifierName(){return "id";}
}
