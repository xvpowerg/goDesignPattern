#  物件導向設計的五原則 SOLID
## 單一職責原則(Single responsibility principle, SRP)
每個物件，不管是類別、函數，負責的功能，都應該只做一件事。
對函數而言，一個函數內，同時做了兩件以上的事情。當發生錯誤時，很難快速定位錯誤的原因。另外，也容易間接導至程式碼的可閱讀性降低。
## 開放封閉原則(Open-Close principle, OCP) 
[程式碼](https://github.com/xvpowerg/goDesignPattern/blob/master/ch1/solid.srp/ocp.go)
 OCP 最高境界要求鄉民們『藉由增加新的程式碼來擴充系統的功能，而不是藉由修改原本已經存在的程式碼來擴充系統的功能』
藉由增加新的程式碼來擴充系統的功能，而不是藉由修改原本已經存在的程式碼來擴充系統的功能。
當需求有異動時，要如何在不變動現在正常運行的程式碼，藉由繼承、相依性注入等方式，增加新的程式碼，以實作新的需求。
假若為了新需求，去修改了原本的程式中的某一個函數，可能會造成其他呼叫使用該函數的的功能，出現非預期的錯誤。
## 里氏替換原則(Liskov substitution principle, LSP)
簡單來說，當實作繼承了 interface 或 base-class的 sub-class，那麼在程式中，只要出現該 interface 或 base-class 的部份，都可以用 sub-class 替換。
## 接口隔離原則(Interface segregation principle, ISP)
針對不同需求的用戶，開放其對應需求的介面，提拱使用。可避免不相關的需求介面異動，造成被強迫一同面對異動的情況。
## 依賴反轉原則(Dependency inversion principle, DIP)
當 A 模組在內部使用 B 模組的情況下，我們稱 A 為高階模組，B 為低階模組。高階模組不應該依賴於低階模組，兩者都該依賴抽象介面。
