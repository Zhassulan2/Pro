package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "teambonussessionperson", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Teambonussessionperson {
    @Id
    @Column(name="Id",unique = true, nullable = false)
    private String id;

    @Column(name = "tbs_id")
    private String tbs_id;
    @Column(name = "person")
    private String person;
    @Column(name = "teambonusforperson")
    private float teambonusforperson;

    public String getIdentifier() {return this.id;}

    public String getIdentifierName(){return "id";}
}
