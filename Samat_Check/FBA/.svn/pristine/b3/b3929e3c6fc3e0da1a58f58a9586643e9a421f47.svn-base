package sync.dto;

import org.hibernate.annotations.OptimisticLockType;
import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "attributesetinstance", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Attributesetinstance implements DataModel{

  @Id
  @Column(name = "id", unique = true, nullable = false)
  private String id;
  @Column(name = "attributeset_Id")
  private String attributeset_Id;
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


  public String getAttributeset_Id() {
    return attributeset_Id;
  }

  public void setAttributeset_Id(String attributeset_Id) {
    this.attributeset_Id = attributeset_Id;
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

  public String getIdentifier() {
    return this.id;
  }
  public String getIdentifierName(){return "id";}
}
