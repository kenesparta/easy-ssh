Easy-SSH for SSH config file management

```
Host key_name
    Hostname 1.2.3.4.5
    User ubuntu
    IdentitiesOnly yes
    IdentityFile ~/.ssh/custom_id_file
    Port 22
```
For the moment, can create the following characteristics:

- Hostname: string
- User: string
- IdentitiesOnly: boolean (yes/no)
- IdentityFile: string
- Port: int
