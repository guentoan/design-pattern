## Factory Method

### Intent

> **Factory Method** is a creational design pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created.

**Factory Method** là một creational design pattern, cung cấp một interface để tạo object cho lớp cha, nhưng cũng cho phép các lớp con có thể sửa đổi đối tượng được tạo.

![Factory Method Intent](https://refactoring.guru/images/patterns/content/factory-method/factory-method-en.png)

### Problem

> Imagine that you’re creating a logistics management application. The first version of your app can only handle transportation by trucks, so the bulk of your code lives inside the `Truck` class.

Bạn hãy tưởng tượng bạn đang tạo một ứng dụng quản lý giao vận. Trong version đầu tiên của ứng dụng chỉ quản lý vận chuyển cho các xe tải, vậy nên phần lớn code của bạn sẽ nằm trong class `Truck`.

> After a while, your app becomes pretty popular. Each day you receive dozens of requests from sea transportation companies to incorporate sea logistics into the app.

Sau đó, ứng dụng của bạn trở nên phổ biến, và bạn nhận được yêu cầu từ các công ty hàng hải để kết hợp với giao vận đường biển.

![Factory Method Problem](https://refactoring.guru/images/patterns/diagrams/factory-method/problem1-en.png)


> Great news, right? But how about the code? At present, most of your code is coupled to the `Truck` class. Adding `Ships` into the app would require making changes to the entire codebase. Moreover, if later you decide to add another type of transportation to the app, you will probably need to make all of these changes again.

Toẹt vời, thế còn code thì sao? Hiện tại hầu hết code của bạn được triển khai ở lớp `Truck`. Việc thêm lớp `Ships` vào trong app sẽ yêu cầu thay đổi với các đối tượng base. Hơn nữa, nếu bạn thêm một phương tiện mới, thì bạn sẽ phải thay đổi code lần nữa.

> As a result, you will end up with pretty nasty code, riddled with conditionals that switch the app’s behavior depending on the class of transportation objects.

Kết quả là bạn sẽ có một mớ code hổ lốn, với rất nhiều điều kiện thay đổi của ứng dụng phụ thuộc vào lớp đối tượng giao vận.

### Solution

> The Factory Method pattern suggests that you replace direct object construction calls (using the `new` operator) with calls to a special *factory* method. Don’t worry: the objects are still created via the `new` operator, but it’s being called from within the factory method. Objects returned by a factory method are often referred to as *products*.

Factory Method pattern gợi ý rằng, thay vì tạo một đối tượng thông qua các lệnh khởi tạo trực tiếp(bằng cách sử dụng toán tử `new`) thì hãy gọi đến phương thức *factory*. Dĩ nhiên là các object vẫn được tạo bằng toán tử `new`, nhưng nó sẽ được gọi từ trong phương thức *factory*. Các object được trả về theo phương thức factory thường được gọi là *products*.

![Factory Method Solution](https://refactoring.guru/images/patterns/diagrams/factory-method/solution1.png)

> At first glance, this change may look pointless: we just moved the constructor call from one part of the program to another. However, consider this: now you can override the factory method in a subclass and change the class of products being created by the method.

Thoạt nhìn, sự thay đổi này có vẻ vô nghĩa: chúng ta chỉ chuyển việc gọi constructor từ phần này sang phần khác. Tuy nhiên, hãy xem xét điều này: bây giờ bạn có thể ghi đè phương thức factory trong một lớp con, và thay đổi lớp product đang được tạo bởi phương thức này.

> There’s a slight limitation though: subclasses may return different types of products only if these products have a common base class or interface. Also, the factory method in the base class should have its return type declared as this interface.

Có một hạn chế nhỏ: các lớp con có thể trả về các kiểu product khác nhau chỉ khi các product này có lớp base hoắc interface chung. Ngoài ra, phương thức factory trong lớp base nên có kiểu trả về là interface đó.

![Factory Method Solution 2](https://refactoring.guru/images/patterns/diagrams/factory-method/solution2-en.png)

> For example, both `Truck` and `Ship` classes should implement the `Transport` interface, which declares a method called `deliver`. Each class implements this method differently: trucks deliver cargo by land, ships deliver cargo by sea. The factory method in the `RoadLogistics` class returns truck objects, whereas the factory method in the `SeaLogistics` class returns ships.

Ví dụ, cả hai lớp `Truck` và `Ship` đều được triển khia từ interface `Transport`, interface này khai báo một method `deliver`. Mỗi lớp sẽ triển khai phương thức này theo cách khác nhau: Truck sẽ deliver hàng hoá theo đường bộ, Ship sẽ deliver hàng hoá bằng đường biển. Phương thức factory `RoadLogistics` sẽ trả về đối tượng `Truck`, còn phương thức factory `SeaLogistics` sẽ trả về đối tượng `Ship`.

![Factory Method Solution 3](https://refactoring.guru/images/patterns/diagrams/factory-method/solution3-en.png)

> The code that uses the factory method (often called the *client code*) doesn’t see a difference between the actual products returned by various subclasses. The client treats all the products as abstract `Transport`. The client knows that all transport objects are supposed to have the `deliver` method, but exactly how it works isn’t important to the client.

Đoạn code sử dụng phương thức factory(thường được gọi là *client code*) không nhìn thấy sự khác biệt giữa những product được trả về bởi các lớp con. Client coi tất cả các product như là lớp trừu tượng `Transport`. Client biết rằng tất cả các đối tượng transport đều có phương thức `deliver`, nhưng nó không quan tâm cách hoạt động chính xác của phương thức đó.

### Structure

![Factory Method Structure](https://refactoring.guru/images/patterns/diagrams/factory-method/structure.png)

> - The **Product** declares the interface, which is common to all objects that can be produced by the creator and its subclasses.

- Interface **Product** được định nghĩa chung cho tất cả các đối tượng, các đối tượng này có thể được tạo bởi creator và lớp con của nó.

> - **Concrete Products** are different implementations of the product interface.

- **Concrete Products** là các triển khai khác nhau của interface product.

> - The **Creator** class declares the factory method that returns new product objects. It’s important that the return type of this method matches the product interface. You can declare the factory method as `abstract` to force all subclasses to implement their own versions of the method. As an alternative, the base factory method can return some default product type.

- Lớp **Creator** định nghĩa phương thức factory, phương thức này trả về một đối tượng product mới. Kiểu trả về của phương thức này phải tương ứng với interface product. Bạn có thể định nghĩa phương thức factory như là `abstract` để tất cả lớp con triển khai theo ý chúng. Và phương thức factory có thể trả về kiểu product mặc định.

> - **Concrete Creators** override the base factory method so it returns a different type of product. Note that the factory method doesn’t have to create new instances all the time. It can also return existing objects from a cache, an object pool, or another source.

- **Concrete Creators** ghi đè phương thức factory base để trả về một kiểu product mới. Lưu ý là không phải lúc nào phương thức factory cũng trả về một đối tượng mới. Nó cũng có thể trả về một đối tượng đã tồn lại từ cache, object pool, hay từ một nguồn nào đó.

### Pseudocode

> This example illustrates how the **Factory Method** can be used for creating cross-platform UI elements without coupling the client code to concrete UI classes.

Ví dụ này minh hoạ cách **Factory Method** được dùng để tạo các phần tử cross-platform UI mà không cần phải gắn code client với lớp UI cụ thể.

![Pseudocode](https://refactoring.guru/images/patterns/diagrams/factory-method/example.png)

> The base `Dialog` class uses different UI elements to render its window. Under various operating systems, these elements may look a little bit different, but they should still behave consistently. A button in Windows is still a button in Linux.

Lớp `Dialog` cở sở dùng để render các phần tử UI khác nhau lên màn hình. Với mỗi hệ điều hành khác nhau, các phần tử này có thể vài khác biệt nhỏ, nhưng chúng vẫn phải đồng nhất. Một button trên windows thì vẫn là một button trên linux.

> When the factory method comes into play, you don’t need to rewrite the logic of the `Dialog` class for each operating system. If we declare a factory method that produces buttons inside the base `Dialog` class, we can later create a subclass that returns Windows-styled buttons from the factory method. The subclass then inherits most of the code from the base class, but, thanks to the factory method, can render Windows-looking buttons on the screen.

Khi sử dụng phương thức factory, bạn không cần phải viết lại logic của lớp `Dialog` cho mỗi hệ điều hành. Nếu chúng ta khai báo phương thức factory để tạo button bên trong lớp `Dialog` base, chúng ta có thể tạo lớp con để trả về button kiểu windowns từ phương thức factory. Lớp con sẽ kế thừa phần lớp code của lớp base, nhưng nhờ vào phương thức factory, chúng ta có thể render các button kiểu windowns lên màn hình.

> For this pattern to work, the base `Dialog` class must work with abstract buttons: a base class or an interface that all concrete buttons follow. This way the code within `Dialog` remains functional, whichever type of buttons it works with.

Với pattern này, lớp `Dialog` base phải làm việc với abstract button(base class hoặc interface) cho tất cả các concrete button. Theo cách này, đoạn code của `Dialog` vẫn hoạt động dù có làm việc với bất kỳ kiểu button nào.

> Of course, you can apply this approach to other UI elements as well. However, with each new factory method you add to the `Dialog`, you get closer to the `Abstract Factory` pattern. Fear not, we’ll talk about this pattern later.

Dĩ nhiên, bạn có thể dùng cách này với bất kỳ phần tử UI nào khác. Tuy nhiên, với mỗi phương thức factory mà bạn thêm vào `Dialog`, bạn sẽ tiến gần hơn tới pattern `Abstract Factory`. Chúng ta sẽ nói về pattern này ở bài viết sau.

```java
// The creator class declares the factory method that must
// return an object of a product class. The creator's subclasses
// usually provide the implementation of this method.
class Dialog is
    // The creator may also provide some default implementation
    // of the factory method.
    abstract method createButton():Button

    // Note that, despite its name, the creator's primary
    // responsibility isn't creating products. It usually
    // contains some core business logic that relies on product
    // objects returned by the factory method. Subclasses can
    // indirectly change that business logic by overriding the
    // factory method and returning a different type of product
    // from it.
    method render() is
        // Call the factory method to create a product object.
        Button okButton = createButton()
        // Now use the product.
        okButton.onClick(closeDialog)
        okButton.render()


// Concrete creators override the factory method to change the
// resulting product's type.
class WindowsDialog extends Dialog is
    method createButton():Button is
        return new WindowsButton()

class WebDialog extends Dialog is
    method createButton():Button is
        return new HTMLButton()


// The product interface declares the operations that all
// concrete products must implement.
interface Button is
    method render()
    method onClick(f)

// Concrete products provide various implementations of the
// product interface.
class WindowsButton implements Button is
    method render(a, b) is
        // Render a button in Windows style.
    method onClick(f) is
        // Bind a native OS click event.

class HTMLButton implements Button is
    method render(a, b) is
        // Return an HTML representation of a button.
    method onClick(f) is
        // Bind a web browser click event.


class Application is
    field dialog: Dialog

    // The application picks a creator's type depending on the
    // current configuration or environment settings.
    method initialize() is
        config = readApplicationConfigFile()

        if (config.OS == "Windows") then
            dialog = new WindowsDialog()
        else if (config.OS == "Web") then
            dialog = new WebDialog()
        else
            throw new Exception("Error! Unknown operating system.")

    // The client code works with an instance of a concrete
    // creator, albeit through its base interface. As long as
    // the client keeps working with the creator via the base
    // interface, you can pass it any creator's subclass.
    method main() is
        this.initialize()
        dialog.render()
```

### Applicability

> **Use the Factory Method when you don’t know beforehand the exact types and dependencies of the objects your code should work with.**

**Sử dụng phương thức factory khi bạn không biết chính xác kiểu phụ thuộc của object mà code của bạn sẽ làm việc.**

> The Factory Method separates product construction code from the code that actually uses the product. Therefore it’s easier to extend the product construction code independently from the rest of the code.

Phương thức factory phân tách code khởi tạo product và code sử dụng product. Do đó, nó sẽ dễ dàng mở rộng việc tạo mới product.

> For example, to add a new product type to the app, you’ll only need to create a new creator subclass and override the factory method in it.

Ví dụ, thêm một loại product mới vào ứng dụng, bạn sẽ chỉ cần tạo một lớp con mới và ghi đè phương thức factory.

> **Use the Factory Method when you want to provide users of your library or framework with a way to extend its internal components.**

**Sử dụng phương thức factory khi bạn muốn cung cấp cho người dùng thư viện hay framework với cách mở rộng các component trong đó.**

> Inheritance is probably the easiest way to extend the default behavior of a library or framework. But how would the framework recognize that your subclass should be used instead of a standard component?

Kế thừa là cách dễ dàng nhất để mở rộng các hành vi mặc định của library hay framework. Nhưng làm thế nào để framework có thể nhận biết rằng nên sử dụng lớp con của bạn thay vì một component tiêu chuẩn?

> The solution is to reduce the code that constructs components across the framework into a single factory method and let anyone override this method in addition to extending the component itself.

Giải pháp là giảm code khởi tạo các component ở framework thành một phương thức factory duy nhất và cho phép bất kỳ ai cũng có thể ghi đè phương thức này để thêm các phần mở rộng cho component đó.

> Let’s see how that would work. Imagine that you write an app using an open source UI framework. Your app should have round buttons, but the framework only provides square ones. You extend the standard `Button` class with a glorious `RoundButton` subclass. But now you need to tell the main `UIFramework` class to use the new button subclass instead of a default one.

Hãy cùng xem cách chúng hoạt động. Hãy tưởng tượng bạn viết một ứng dụng sử dụng một framework UI mã nguồn mở. Ứng dụng của bạn nên có một button góc tròn, nhưng framework chỉ cung cấp button góc vuông. Bạn cần phải mở rộng lớp `Button` tiêu chuẩn thành lớp con `RoundButton`. Và bạn cần nói cho lớp `UIFramework` rằng bạn sử dụng lớp con button mới thay vì lớp mặc định.

> To achieve this, you create a subclass `UIWithRoundButtons` from a base framework class and override its `createButton` method. While this method returns `Button` objects in the base class, you make your subclass return `RoundButton` objects. Now use the `UIWithRoundButtons` class instead of `UIFramework`. And that’s about it!

Để làm được điều đó, bạn cần tạo lớp con `UIWithRoundButtons` từ lớp base của framework, và ghi đè phương thức `createButton`. Trong khi phương thức này trả về object `Button`, thì lớp con của bạn cần trả về object `RoundButton`. Bây giờ bạn sử dụng `UIWithRoundButtons` thay vì `UIFramework`. And that’s about it!

> **Use the Factory Method when you want to save system resources by reusing existing objects instead of rebuilding them each time.**

**Sử dụng phương thức factory khi bạn muốn tiết kiệm tài nguyên bằng cách tái sử dụng object thay vì tạo mới chúng mỗi khi sử dụng.**

> You often experience this need when dealing with large, resource-intensive objects such as database connections, file systems, and network resources.

Bạn thường gặp vấn đề này khi làm việc với các đối tượng lớn, sử dụng nhiều tài nguyên như là kết nối database, file, và network.

> Let’s think about what has to be done to reuse an existing object:
> 1. First, you need to create some storage to keep track of all of the created objects.
> 2. When someone requests an object, the program should look for a free object inside that pool.
> 3. … and then return it to the client code.
> 4. If there are no free objects, the program should create a new one (and add it to the pool).

Hãy thử nghĩ về những gì phải làm để tái sử dụng một đối tượng hiện có.

1. Thứ nhất, bạn cần tạo một nơi lưu trữ tất cả các đối tượng đã tạo.
2. Khi ai đó yêu cầu một đối tượng, chương trình sẽ tìm kiếm đối tượng đó trong pool.
3. ... và sau đó trả về cho client.
4. Nếu không có, chương trình sẽ tạo mới một đối tương và thêm nó vào pool.

> That’s a lot of code! And it must all be put into a single place so that you don’t pollute the program with duplicate code.

Có khá nhiều code! Và nó phải được đặt vào một nơi để tránh cách đoạn mã trùng lặp.

> Probably the most obvious and convenient place where this code could be placed is the constructor of the class whose objects we’re trying to reuse. However, a constructor must always return **new objects** by definition. It can’t return existing instances.

Có lẽ nơi rõ ràng và thuận tiện nhất là hàm khởi tạo của class mà chúng ta đang cố gắng tái sử dụng. Tuy nhiên, một hàm khởi tạo luôn phải trả về một **new objects**. Nó không thể trả lại các instance đang có.

> Therefore, you need to have a regular method capable of creating new objects as well as reusing existing ones. That sounds very much like a factory method.

Do đó, bạn cần phải có một phương thức có khả năng tạo mới một object cũng như tái sử dụng object đã tồn tại. Vâng, đó chính là phương thức factory.

### How to Implement

> 1. Make all products follow the same interface. This interface should declare methods that make sense in every product.

1. Tạo tất cả các product với cùng một interface. Interface này nên định nghĩa các phương thức có ý nghĩa với tất cả các product.

> 2. Add an empty factory method inside the creator class. The return type of the method should match the common product interface.

2. Thêm một phương thức factory rỗng bên trong lớp creator. Phương thức này trả về kiểu của interface product.

> 3. In the creator’s code find all references to product constructors. One by one, replace them with calls to the factory method, while extracting the product creation code into the factory method.
You might need to add a temporary parameter to the factory method to control the type of returned product.
>At this point, the code of the factory method may look pretty ugly. It may have a large `switch` statement that picks which product class to instantiate. But don’t worry, we’ll fix it soon enough.

3. Trong creator, tìm tất cả các tham chiếu đến khởi tạo product. Thay thế toàn bộ lệnh khởi tạo bằng phương thức factory, và đưa lệnh khởi tạo đó vào bên trong phương thức factory. Bạn có thể cần thêm tham số vào phương thức factory để điều khiển kiểu product trả về.
Ở bước này, code của phương thức factory nhìn không được đẹp lắm. Bạn có thể có một đoạn code `switch` lớn để lựa chọn loại product khởi tạo. Nhưng đừng lo, chúng ta sẽ khắc phục điều này sớm thôi.

> 4. Now, create a set of creator subclasses for each type of product listed in the factory method. Override the factory method in the subclasses and extract the appropriate bits of construction code from the base method.

4. Bây giờ, tạo một tập các lớp con của creator cho mỗi kiểu product trong phương thức factory. Ghi đè phương thức factory ở lớp con.

> 5. If there are too many product types and it doesn’t make sense to create subclasses for all of them, you can reuse the control parameter from the base class in subclasses.
> For instance, imagine that you have the following hierarchy of classes: the base `Mail` class with a couple of subclasses: `AirMail` and `GroundMail`; the `Transport` classes are `Plane`, `Truck` and `Train`. While the `AirMail` class only uses `Plane` objects, `GroundMail` may work with both `Truck` and `Train` objects. You can create a new subclass (say `TrainMail`) to handle both cases, but there’s another option. The client code can pass an argument to the factory method of the `GroundMail` class to control which product it wants to receive.

5. Nếu có qúa nhiều kiểu product và nó không phù hợp tạo lớp con cho chúng, bạn có thể tái sử dụng tham số đầu vào từ lớp base.
Ví dụ: hãy tưởng tượng bạn có một list các class kế thừa như sau: Lớp base `Mail` và lớp con: `AirMail` và `GroundMail`; lớp `Transport` và lớp con: `Plane`, `Truck` và `Train`. Trong khi lớp `AirMail` chỉ sử dụng đối tượng `Plant`, `GroundMail` có thể sử dụng object `Truck` và `Train`. Bạn có thể tạo lớp con (gọi là `TrainMail`) để xử lý cả hai trường hợp trên, nhưng cũng có các lựa chọn khác. Client có thể gửi tham số vào phương thức factory của `GroundMail` để điều khiển kiểu product mà nó muốn nhận.

> 6. If, after all of the extractions, the base factory method has become empty, you can make it abstract. If there’s something left, you can make it a default behavior of the method.

6. Nếu, sau cùng phương thức factory của lớp base mà rỗng, thì bạn có thể chuyển nó thành lớp abstract. Nếu còn thứ gì đó thì bạn có thể coi đó là hành vi mặc định của phương thức.


### Pros and Cons

#### Pros

>- You avoid tight coupling between the creator and the concrete products.

- Tránh được kết hợp quá chặt chẽ giữa creator và concrete product.

>- Single Responsibility Principle. You can move the product creation code into one place in the program, making the code easier to support.

- Single Responsibility Principle. Bạn có thể di chuyển code tạo product vào một nơi trong chương trình, giúp code hỗ trợ dễ dàng hơn.

>- Open/Closed Principle. You can introduce new types of products into the program without breaking existing client code.

- Open/Closed Principle. Bạn có thể thêm các kiểu product mới vào chương trình, mà không làm ảnh hưởng đến code client hiện tại.

#### Cons
> - The code may become more complicated since you need to introduce a lot of new subclasses to implement the pattern. The best case scenario is when you’re introducing the pattern into an existing hierarchy of creator classes.

- Code của bạn có thể trở nên phức tạp khi bạn thêm vào quá nhiều lớp con để triển khai pattern. Trường hợp tốt nhất là khi bạn triển khai pattern bằng cách sử dụng kế thừa của lớp creator.

### Relations with Other Patterns

> - Many designs start by using **Factory Method** (less complicated and more customizable via subclasses) and evolve toward **Abstract Factory**, **Prototype**, or **Builder** (more flexible, but more complicated).

- Nhiều pattern bắt đầu bằng cách sử dụng **Factory Method** (ít phức tạp hơn và có thể tùy chỉnh nhiều hơn thông qua các lớp con) và phát triển theo hướng **Abstract Factory**, **Prototype** hoặc **Builder** (linh hoạt hơn nhưng phức tạp hơn).

> - **Abstract Factory** classes are often based on a set of **Factory Methods**, but you can also use **Prototype** to compose the methods on these classes.

- Các lớp **Abstract Factory** thường dựa trên một tập các **Factory Method**, nhưng bạn cũng có thể sử dụng **Prototype** để cấu trúc các phương thức trên các lớp này.

> - You can use **Factory Method** along with **Iterator** to let collection subclasses return different types of iterators that are compatible with the collections.

- Bạn có thể sử dụng **Factory Method** cùng với **Iterator** để cho phép các lớp con của collection trả về các kiểu vòng lặp khác nhau tương thích với các collection.

> - **Prototype** isn’t based on inheritance, so it doesn’t have its drawbacks. On the other hand, **Prototype** requires a complicated initialization of the cloned object. **Factory Method** is based on inheritance but doesn’t require an initialization step.

- **Prototype** không dựa trên sự kế thừa, vì vậy nó không có nhược điểm. Mặt khác, **Prototype** yêu cầu khởi tạo nhân bản đối tượng phức tạp. **Factory Method** dựa trên kế thừa nhưng không yêu cầu bước khởi tạo.

> - **Factory Method** is a specialization of **Template Method**. At the same time, a **Factory Method** may serve as a step in a large **Template Method**.

- **Factory Method** là một chuyên môn hóa của **Template Method**. Đồng thời, **Factory Method** có thể đóng vai trò là một bước trong một **Template Method** lớn.

### Code Examples

- [C#](https://refactoring.guru/design-patterns/factory-method/csharp/example)
- [C++](https://refactoring.guru/design-patterns/factory-method/cpp/example)
- [Golang](https://refactoring.guru/design-patterns/factory-method/go/example)
- [Java](https://refactoring.guru/design-patterns/factory-method/java/example)
- [PHP](https://refactoring.guru/design-patterns/factory-method/php/example)
- [Python](https://refactoring.guru/design-patterns/factory-method/python/example)
- [Ruby](https://refactoring.guru/design-patterns/factory-method/ruby/example)
- [Rust](https://refactoring.guru/design-patterns/factory-method/rust/example)
- [Swift](https://refactoring.guru/design-patterns/factory-method/swift/example)
- [Typescript](https://refactoring.guru/design-patterns/factory-method/typescript/example)
