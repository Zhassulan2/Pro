package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "categories", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Categories implements DataModel{

  @Id
  @Column(name = "id", unique = true, nullable = false)
  private String id;
  @Column(name = "name")
  private String name;
  @Column(name = "parentid")
  private String parentid;
  @Column(name = "image")
  private String image;
  //@Column(name = "sync")
  //private long sync;


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


  public String getParentid() {
    return parentid;
  }

  public void setParentid(String parentid) {
    this.parentid = parentid;
  }


  public String getImage() {
    return image;
  }

  public void setImage(String image) {
    this.image = image;
  }

/*
  public long getSync() {
    return sync;
  }

  public void setSync(long sync) {
    this.sync = sync;
  }
*/
  public String getIdentifier() {
    return this.id;
  }
  public String getIdentifierName(){return "id";}
}
