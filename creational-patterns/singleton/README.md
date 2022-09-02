# Singleton

## Intent

> **Singleton** is a creational design pattern that lets you ensure that a class has only one instance, while providing a global access point to this instance.

**Singleton** là một design pattern thuộc nhóm creational, giúp bạn tạo ra một lớp chỉ với một instance duy nhất, trong khi cung cấp điểm truy cập toàn cục cho instance đấy.

![singleton-intent](https://refactoring.guru/images/patterns/content/singleton/singleton.png)

## Problem

> The Singleton pattern solves two problems at the same time, violating the *Single Responsibility Principle*:

Singleton cùng lúc giải quyết hai vấn đề, vi phạm đến *nguyên tắc Single Responsibility*:

> 1 **Ensure that a class has just a single instance.** Why would anyone want to control how many instances a class has? The most common reason for this is to control access to some shared resource—for example, a database or a file.

1 **Đảm bảo mỗi lớp chỉ có một instance**. Tại sao ai cũng muốn kiểm soát số lượng instance của một class? Lý do phổ biến nhất là để quản lý truy cập đến các tài nguyên chung - ví dụ như truy cập vào cơ sở dữ liệu hay file.

> Here’s how it works: imagine that you created an object, but after a while decided to create a new one. Instead of receiving a fresh object, you’ll get the one you already created.

Đây là cách hoạt động của singletone. Hãy tưởng tượng, bạn đã tạo một đối tượng nhưng sau một thời gian bạn quyết định tạo một đối tượng mới. Bởi vì Singleton bảo đảm chỉ có duy nhất một instance được tạo ra, nên thay vì nhận một đối tượng hoàn toàn mới, bạn sẽ nhận về đối tượng đã được tạo trước đó.

> Note that this behavior is impossible to implement with a regular constructor since a constructor call **must** always return a new object by design.

Lưu ý rằng cách này không thể thực thi với các hàm constructor thông thường vì chúng sẽ luôn trả về đối tượng mới.

![singleton-problem](https://refactoring.guru/images/patterns/content/singleton/singleton-comic-1-en.png)

> 2 **Provide a global access point to that instance**. Remember those global variables that you (all right, me) used to store some essential objects? While they’re very handy, they’re also very unsafe since any code can potentially overwrite the contents of those variables and crash the app.

2 **Cung cấp một điểm truy cập toàn cục cho các instance**. Hãy nhớ lại các biến toàn cục mà bạn đã dùng để lưu trữ một số đối tượng thiết yếu. Mặc dù khá tiện dụng, nhưng chúng cũng không an toàn vì bất kỳ đoạn code nào cũng có khả năng ghi đè lên nội dung của biến đó và làm chương trình bị crash.

> Just like a global variable, the Singleton pattern lets you access some object from anywhere in the program. However, it also protects that instance from being overwritten by other code.

Giống như biến toàn cục, Singleton giúp bạn truy cập đến các đối tượng ở bất cứ đâu trong chương trình. Tuy nhiên, không như biến, nó cũng sẽ bảo vệ các thực thể khỏi bị ghi đè bởi code khác.

> There’s another side to this problem: you don’t want the code that solves problem #1 to be scattered all over your program. It’s much better to have it within one class, especially if the rest of your code already depends on it.

Mặt khác: bạn không muốn code giải quyết vấn đề #1 bị phân tán khắp chương trình. Nó sẽ tốt hơn nếu có nó trong một lớp, đặc biệt nếu phần còn lại của code bạn đã phụ thuộc vào nó.

> Nowadays, the Singleton pattern has become so popular that people may call something a *singleton* even if it solves just one of the listed problems.

Ngày nay, Singleton đã trở nên phổ biến đến mức mọi người có thể gọi một thứ gì đó là singleton, ngay cả khi nó chỉ giải quyết được một trong những vấn đề được liệt kê ở trên.

## Solution

> All implementations of the Singleton have these two steps in common:

Tất cả các triển khai của singleton đều có hai bước là:

> - Make the default constructor private, to prevent other objects from using the new operator with the Singleton class.

 - các hàm constructor là private, để tránh các đối tượng khác sử dụng từ khoá `new` với lớp singleton.

> - Create a static creation method that acts as a constructor. Under the hood, this method calls the private constructor to create an object and saves it in a static field. All following calls to this method return the cached object.

 - Tạo một hàm static thay thế cho hàm constructor. Bên trong hàm này, sẽ gọi đến hàm constructor để tạo một object và lưu nó vào một biến static. Tất cả những lệnh gọi hàm này sẽ trả về một object cache.

> If your code has access to the Singleton class, then it’s able to call the Singleton’s static method. So whenever that method is called, the same object is always returned.

Nếu code của bạn có quyền truy cập vào lớp singleton thì nó có thể gọi hàm static của singleton. Vì vậy, bất cứ hàm nào được gọi, thì sẽ có cùng một đối tượng trả về.

## Real-World Analogy

> The government is an excellent example of the Singleton pattern. A country can have only one official government. Regardless of the personal identities of the individuals who form governments, the title, “The Government of X”, is a global point of access that identifies the group of people in charge.

Trong thế giới thực, chính phủ là một ví dụ điển hình của singleton. Mỗi quốc gia chỉ có duy nhất một chính phủ chính thức. Bất kể danh tính cá nhân của người đại diện của chính phủ, chức danh, "Chính phủ X", đều là điểm truy cập để xác định danh tính của nhóm người phụ trách.

## Structure

![singleton-structure](https://refactoring.guru/images/patterns/diagrams/singleton/structure-en.png)

> The **Singleton** class declares the static method getInstance that returns the same instance of its own class.
> The Singleton’s constructor should be hidden from the client code. Calling the getInstance method should be the only way of getting the Singleton object.

Lớp singleton sẽ định nghĩa hàm static `getInstance` để trả về cùng một instance của lớp đó.

Hàm constructor của singleton sẽ bị ẩn ở client. Cách duy nhất để tạo đối tượng singleton là gọi đến hàm `getInstance`.

## Pseudocode

> In this example, the database connection class acts as a **Singleton**. This class doesn’t have a public constructor, so the only way to get its object is to call the getInstance method. This method caches the first created object and returns it in all subsequent calls.

Trong ví dụ này, lớp kết nối database đóng vai trò là một **singleton**. Lớp này không có hàm constructor public, cách duy nhất để tạo đối tượng này là gọi đến hàm `getInstance`. Hàm này sẽ khởi tạo đối tượng đầu tiên, rồi cache chúng lại, và sử dụng nó cho lần gọi tiếp theo.

```java
// The Database class defines the `getInstance` method that lets
// clients access the same instance of a database connection
// throughout the program.
class Database is
    // The field for storing the singleton instance should be
    // declared static.
    private static field instance: Database

    // The singleton's constructor should always be private to
    // prevent direct construction calls with the `new`
    // operator.
    private constructor Database() is
        // Some initialization code, such as the actual
        // connection to a database server.
        // ...

    // The static method that controls access to the singleton
    // instance.
    public static method getInstance() is
        if (Database.instance == null) then
            acquireThreadLock() and then
                // Ensure that the instance hasn't yet been
                // initialized by another thread while this one
                // has been waiting for the lock's release.
                if (Database.instance == null) then
                    Database.instance = new Database()
        return Database.instance

    // Finally, any singleton should define some business logic
    // which can be executed on its instance.
    public method query(sql) is
        // For instance, all database queries of an app go
        // through this method. Therefore, you can place
        // throttling or caching logic here.
        // ...

class Application is
    method main() is
        Database foo = Database.getInstance()
        foo.query("SELECT ...")
        // ...
        Database bar = Database.getInstance()
        bar.query("SELECT ...")
        // The variable `bar` will contain the same object as
        // the variable `foo`.
```

## Applicability

>Use the Singleton pattern when a class in your program should have just a single instance available to all clients; for example, a single database object shared by different parts of the program.

**Sử dụng singleton khi một lớp trong chương trình của bạn chỉ nên có một instance duy nhất cho tất cả các clients; ví dụ, đối tượng database được sử dụng bởi nhiều nơi khác nhau của chương trình.**

> The Singleton pattern disables all other means of creating objects of a class except for the special creation method. This method either creates a new object or returns an existing one if it has already been created.

Singleton vô hiệu hoá tất cả các hàm tạo đối tượng của lớp trừ hàm tạo đặc biệt. Hàm này tạo một đối tượng mới hoặc trả về một đối tượng đã có nếu nó đã được tạo trước đó.

> Use the Singleton pattern when you need stricter control over global variables.

**Sử dụng Singleton khi bạn cần kiểm soát chặt chẽ hơn đối với các biến toàn cục.**

> Unlike global variables, the Singleton pattern guarantees that there’s just one instance of a class. Nothing, except for the Singleton class itself, can replace the cached instance.

Không giống như các biến toàn cục, singleton đảm bảo rằng chỉ có duy nhất một instance. không có gì khác, ngoài instance của chính nó, có thể thay đổi được instance được cache.

> Note that you can always adjust this limitation and allow creating any number of Singleton instances. The only piece of code that needs changing is the body of the getInstance method.

Lưu ý rằng bạn có thể điều chỉnh giới hạn và cho phép singleton tạo nhiều instance. Đoạn code cần thay đổi duy nhất là phần thân của hàm `getInstance`.

## How to Implement

> 1. Add a private static field to the class for storing the singleton instance.

1. Thêm một biến static vào trong class lưu instance singleton

> 2. Declare a public static creation method for getting the singleton instance.

2. Khai báo phương thức public static để tạo signleton instance.

> 3. Implement “lazy initialization” inside the static method. It should create a new object on its first call and put it into the static field. The method should always return that instance on all subsequent calls.

3. Triển khai `lazy initialization` bên trong hàm static. Nó tạo ra đối tượng mới cho lần gọi đầu tiên, và put vào biến static. Hàm này luôn trả về instance đó cho tất cả lần gọi tiếp theo.

> 4. Make the constructor of the class private. The static method of the class will still be able to call the constructor, but not the other objects.

4. Tạo hàm construcor private. Hàm static của lớp vẫn có thể gọi đến hàm constructor, nhưng các đối tượng khác thì không.

> 5. Go over the client code and replace all direct calls to the singleton’s constructor with calls to its static creation method.

5. Ở phía client cần thay thế tất cả các cách tạo instance trực tiếp bằng cách gọi đến hàm khởi tạo static.

## Pros and Cons

### Ưu điểm
> - You can be sure that a class has only a single instance.
> - You gain a global access point to that instance.
> - The singleton object is initialized only when it’s requested for the first time.

- Bạn có thể chắc chắn mỗi lớp chỉ có duy nhất một instance.
- Bạn có điểm truy cập toàn cục đến instance đó.
- Đối tượng singleton chỉ được tạo một lần duy nhất trong lần gọi đầu tiên.

### Nhược điểm 

> - Violates the Single Responsibility Principle. The pattern solves two problems at the time.

 - Cả hai vấn đề của singleton đều vi phạm nguyên tắc *Single Responsibility Principle*.

> - The Singleton pattern can mask bad design, for instance, when the components of the program know too much about each other.

 - Singleton có thể che dấu các thiết kế tồi, ví dụ một instance, khi các thành phần của chương trình biết quá nhiều về nhau.

> - The pattern requires special treatment in a multithreaded environment so that multiple threads won’t create a singleton object several times.

 - Pattern này yêu cầu phải xử lý cẩn thận trong đa luồng, để nhiều luồng không thể khởi tạo object nhiều lần.

> - It may be difficult to unit test the client code of the Singleton because many test frameworks rely on inheritance when producing mock objects. Since the constructor of the singleton class is private and overriding static methods is impossible in most languages, you will need to think of a creative way to mock the singleton. Or just don’t write the tests. Or don’t use the Singleton pattern.

 - Có thể gặp khó khăn khi thực hiện unit test cho code client của Singleton, vì các framework test dựa trên kế thừa khi tạo ra các mock objects. Và hàm khởi tạo của lớp Singleton là private cùng với việc ghi đè các phương thức tĩnh là không thể trong hầu hết các ngôn ngữ, nên bạn sẽ cần phải nghĩ ra một cách sáng tạo để mô phỏng lớp singleton. Hoặc không thực hiện test. Hoặc không sử dụng Singleton.

## Relations with Other Patterns

> - A Facade class can often be transformed into a Singleton since a single facade object is sufficient in most cases.

 - Một lớp **Facade** thường có thể được chuyển đổi thành Singleton vì một đối tượng facade duy nhất là đủ trong hầu hết các trường hợp.

> - Flyweight would resemble Singleton if you somehow managed to reduce all shared states of the objects to just one flyweight object. But there are two fundamental differences between these patterns:
>   - There should be only one Singleton instance, whereas a Flyweight class can have multiple instances with different intrinsic states.
>   - The Singleton object can be mutable. Flyweight objects are immutable.

 - **Flyweight** sẽ giống với Singleton nếu bạn bằng cách nào đó giảm được tất cả các trạng thái được chia sẻ của các đối tượng xuống chỉ còn một đối tượng flyweight. Nhưng có hai điểm khác biệt cơ bản giữa các pattern này:
   - Chỉ nên có một instance Singleton, trong khi lớp Flyweight có thể có nhiều instance với các trạng thái nội tại khác nhau.
   - Đối tượng singleton có thể thay đổi được, đối tượng flyweight là không đổi.

> - Abstract Factories, Builders and Prototypes can all be implemented as Singletons.

 - **Abstract Factory**, **Builder** và **Prototype** đều có thể được triển khai dưới dạng các Singletons.

## Code Examples

 - [Java](https://refactoring.guru/design-patterns/singleton/java/example)
 - [C#](https://refactoring.guru/design-patterns/singleton/csharp/example)
 - [C++](https://refactoring.guru/design-patterns/singleton/cpp/example)
 - [PHP](https://refactoring.guru/design-patterns/singleton/php/example)
 - [Python](https://refactoring.guru/design-patterns/singleton/python/example)
 - [Ruby](https://refactoring.guru/design-patterns/singleton/ruby/example)
 - [Swift](https://refactoring.guru/design-patterns/singleton/swift/example)
 - [Typescript](https://refactoring.guru/design-patterns/singleton/typescript/example)
 - [Golang](https://refactoring.guru/design-patterns/singleton/go/example)