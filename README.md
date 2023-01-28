# GO-MAC-Changer
Bu kod, ifconfig komutunu kullanarak belirtilen ağ arayüzünün (interfaceName değişkeninde belirtilen) 
MAC adresini belirtilen yeni MAC adresine (newMAC değişkeninde belirtilen) değiştirir. 
Bu işlem için "sudo" komutu ile sistem yönetici yetkileri kullanılıyor. 
exec paketi kullanılarak os/exec kütüphanesi kullanılıyor. Eğer herhangi bir hata oluşursa, panic() fonksiyonu ile program sonlandırılıyor.
