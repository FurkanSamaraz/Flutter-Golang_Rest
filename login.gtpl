<html>
    <head>
     <title>Kayıt</title>
    <link rel="stylesheet" href="/common.css">
    </head>
    <body>
        <h1>Merhaba</h1>
    <p>Lütfen Kullanıcı Adı ve  Şifre ve numara giriniz. <br/>
   
  
    
    
        <form action="/login" method="post">
            Kullanıcı Adı:<br/><input type="json" name="username"><br/>
          <p> Kullanıcı Soyadı:<br/><input type="json" name="password"><br/><br/>
      
           
               <input type="button" onclick="window.location.href = '/token';" value="ok"/><br/><br/><br/>
       <input type="button" onclick="window.location.href = '/';" value="Geri"/><br/><br/><br/>
                 Teşekkürler...</p>
        </form>
    </body>
</html>