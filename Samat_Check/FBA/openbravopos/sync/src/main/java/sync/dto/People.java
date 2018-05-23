package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "people", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class People implements DataModel{

  @Id
  @Column(name = "id", unique = true, nullable = false)
  private String id;
  @Column(name = "name")
  private String name;
  @Column(name = "apppassword")
  private String apppassword;
  @Column(name = "card")
  private String card;
  @Column(name = "role")
  private String role;
  @Column(name = "visible")
  private String visible;
  @Column(name = "image")
  private String image;
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


  public String getApppassword() {
    return apppassword;
  }

  public void setApppassword(String apppassword) {
    this.apppassword = apppassword;
  }


  public String getCard() {
    return card;
  }

  public void setCard(String card) {
    this.card = card;
  }


  public String getRole() {
    return role;
  }

  public void setRole(String role) {
    this.role = role;
  }


  public String getVisible() {
    return visible;
  }

  public void setVisible(String visible) {
    this.visible = visible;
  }


  public String getImage() {
    return image;
  }

  public void setImage(String image) {
    this.image = image;
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
