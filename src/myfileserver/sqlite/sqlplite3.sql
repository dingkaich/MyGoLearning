create table fileinfo (
    filename varchar(1024) not null,
    filemd5  varchar(512),

    fileserverpath varchar(1024),
    serverIp    varchar(25),
    clientIp varchar(25),
    
)
