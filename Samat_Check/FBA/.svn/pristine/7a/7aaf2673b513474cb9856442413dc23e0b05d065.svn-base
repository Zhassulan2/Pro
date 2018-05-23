package sync.dto;


import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "teambonusconfighosts", uniqueConstraints = {@UniqueConstraint(columnNames = "host")})
public class Teambonusconfighosts implements DataModel{
    @Id
    @Column(name="host",unique = true, nullable = false)
    private String host;

    public String getHost() {
        return host;
    }

    public void setHost(String host) {
        this.host = host;
    }

    public String getIdentifier() {return this.host;}

    public String getIdentifierName(){return "host";}
}
