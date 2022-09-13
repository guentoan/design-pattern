# Prototype

## Intent

> **Prototype** is a creational design pattern that lets you copy existing objects without making your code dependent on their classes.

**Prototype** là một design pattern thuộc nhóm creational, nó giúp bạn sao chép một object mà code của bạn sẽ không phụ thuộc vào lớp của object đó.

![prototype](https://refactoring.guru/images/patterns/content/prototype/prototype.png)

## Problem

> Say you have an object, and you want to create an exact copy of it. How would you do it? First, you have to create a new object of the same class. Then you have to go through all the fields of the original object and copy their values over to the new object.

Giả sử bạn đang có một object, và bạn muốn tạo ra bản sao của nó. Vậy làm thế nào? Đầu tiên, bạn sẽ tạo ra một object mới có cùng lớp, sau đó bạn sẽ lấy giá trị từ tất cả các trường của object gốc và gán nó sang cho object mới.

> Nice! But there’s a catch. Not all objects can be copied that way because some of the object’s fields may be private and not visible from outside of the object itself.

Tuyệt vời! Nhưng có một vấn đề. Không phải tất cả các object đều có thể copied theo cách này, vì có thể có một vài trường là private và không thể truy cập từ bên ngoài.

![](https://refactoring.guru/images/patterns/content/prototype/prototype-comic-1-en.png)
*Copying an object “from the outside” **isn’t** always possible.*

> There’s one more problem with the direct approach. Since you have to know the object’s class to create a duplicate, your code becomes dependent on that class. If the extra dependency doesn’t scare you, there’s another catch. Sometimes you only know the interface that the object follows, but not its concrete class, when, for example, a parameter in a method accepts any objects that follow some interface.

Có nhiều vấn đề với cách tiếp cận này.
 - Khi bạn biết class của object mà bạn tạo bản sao, bạn sẽ trở nên phụ thuộc vào class đó.
 - Thỉnh thoảng bạn chỉ biết interface của object chứ không biết class cụ thể. Khi đó, tham số trong method của bạn sẽ chấp nhận bất kỳ object nào có interface đó.

## Solution

> The Prototype pattern delegates the cloning process to the actual objects that are being cloned. The pattern declares a common interface for all objects that support cloning. This interface lets you clone an object without coupling your code to the class of that object. Usually, such an interface contains just a single `clone` method.

Pattern prototype giao nhiệm vụ sao chép cho các object thực đang được sao chép.(Chỗ này nghĩa là hàm clone sẽ được implement bên trong class của object được sao chép) Pattern này định nghĩa một interface chung hỗ trợ việc sao chép cho tất cả các đối tượng. Interface này giúp bạn sap chép một object mà không cần quan tâm đến class của object đó. Thông thường, interface đó chỉ chứa một method `clone`.

> The implementation of the `clone` method is very similar in all classes. The method creates an object of the current class and carries over all of the field values of the old object into the new one. You can even copy private fields because most programming languages let objects access private fields of other objects that belong to the same class.

Việc triển khia hàm `clone` là tương tự nhau trong mọi class. Phương thức này tạo một đối tượng của class hiện tại và sao chép tất cả giá trị của các biến của object cũ sang object mới. Bạn cũng có thể sao chép các biến private vì phần lớn ngôn ngữ lập trình đều cho phép bạn truy cập vào biến private của object khác có cùng class.

> An object that supports cloning is called a *prototype*. When your objects have dozens of fields and hundreds of possible configurations, cloning them might serve as an alternative to subclassing.

Một object hỗ trợ việc sao chép gọi là *prototype*. Khi object của bạn có hàng chục, hàng trăm biến thì việc sao chép chúng được xem như một giải phải pháp thay thế cho lớp con.

![](https://refactoring.guru/images/patterns/content/prototype/prototype-comic-2-en.png)

> Here’s how it works: you create a set of objects, configured in various ways. When you need an object like the one you’ve configured, you just clone a prototype instead of constructing a new object from scratch.

Vậy nó hoạt động như nào? Bạn tạo một tập các object, được cấu hình theo nhiều cách khác nhau. Khi bạn muốn một object giống như object mà bạn đã cấu hình trước đó, bạn chỉ cần sao chép một prototype thay vì tạo một đối tượng từ đầu.

## Real-World Analogy

> In real life, prototypes are used for performing various tests before starting mass production of a product. However, in this case, prototypes don’t participate in any actual production, playing a passive role instead.

Trong thực tế, prototypes được sử dụng để thực hiện các thử nghiệm trước khi bắt đầu sản xuất hàng loạt một sản phẩm. Tuy nhiên, trong trường hợp này, prototypes không tham gia vào bất kỳ một quá trình sản xuất thực nào, mà nó chỉ đóng vai trò thụ động trong quá trình sao chép.(được sao chép bằng máy móc chứ sản phẩm không tự sao chép ra)

![](https://refactoring.guru/images/patterns/content/prototype/prototype-comic-3-en.png)

> Since industrial prototypes don’t really copy themselves, a much closer analogy to the pattern is the process of mitotic cell division (biology, remember?). After mitotic division, a pair of identical cells is formed. The original cell acts as a prototype and takes an active role in creating the copy.

Các prototype trong công nghiệp không thực sự tự sao chép, có một ví dụ gần giống với pattern prototype hơn đó là quá trình phân bào trong sinh học. Sau khi phân bào, một cặp tế bào giống hệt nhau được hình thành. Tế bào gốc lúc này đóng vai trò như một prototype và chủ động trong quá trình sao chép.

## Structure

### Basic implementation

![](https://refactoring.guru/images/patterns/diagrams/prototype/structure.png)

> - The **Prototype** interface declares the cloning methods. In most cases, it’s a single clone method.
- **Prototype** inaterface định nghĩa method clone. Trong hầu hết các trường hợp, nó chỉ có một method `clone`.
> - The **Concrete Prototype** class implements the cloning method. In addition to copying the original object’s data to the clone, this method may also handle some edge cases of the cloning process related to cloning linked objects, untangling recursive dependencies, etc.
- Lớp **Concrete Prototype** triển khai phương thức `clone`. Ngoài việc sao chép dữ liệu của đối tượng gốc thì phương thức này cũng có thể xử lý một số vấn đề trong quá trình sao chép như linked objects, hoặc đệ quy...
> - The **Client** can produce a copy of any object that follows the prototype interface.
- **Client** có thể tạo một bản sao bất kỳ đối tượng nào triển khai interface prototype.


### Prototype registry implementation

![](https://refactoring.guru/images/patterns/diagrams/prototype/structure-prototype-cache.png)

> The **Prototype Registry** provides an easy way to access frequently-used prototypes. It stores a set of pre-built objects that are ready to be copied. The simplest prototype registry is a `name → prototype` hash map. However, if you need better search criteria than a simple name, you can build a much more robust version of the registry.

**Prototype Registry** cung cấp cách truy cập dễ dàng tới các prototype được sử dụng thường xuyên. Nó lưu trữ một tập các object đã pre-build để phục vụ việc sao chép. Một prôttype registry đơn giản nhất là một hashmap `name -> prototype`. Tuy nhiên, Nếu bạn cần nhiều từ khoá tìm kiếm hơn là name, thì bạn có thể tự build một registry mạnh mẽ hơn.

## Pseudocode

> In this example, the Prototype pattern lets you produce exact copies of geometric objects, without coupling the code to their classes.

Trong ví dụ này, pattern prototype sẽ chps phép bạn tạo các bản sao của các đối tượng hình học mà không vần phải phụ thuộc vào code của class.

![](https://refactoring.guru/images/patterns/diagrams/prototype/example.png)

> All shape classes follow the same interface, which provides a cloning method. A subclass may call the parent’s cloning method before copying its own field values to the resulting object.

Tất cả các lớp shape đều implement một interface có hàm clone. Trong lớp con có thể gọi phương thức clone của lớp cha trước khi sao chép các biến của chính nó vào đối tượng trả về.

```java
// Base prototype.
abstract class Shape is
    field X: int
    field Y: int
    field color: string

    // A regular constructor.
    constructor Shape() is
        // ...

    // The prototype constructor. A fresh object is initialized
    // with values from the existing object.
    constructor Shape(source: Shape) is
        this()
        this.X = source.X
        this.Y = source.Y
        this.color = source.color

    // The clone operation returns one of the Shape subclasses.
    abstract method clone():Shape


// Concrete prototype. The cloning method creates a new object
// in one go by calling the constructor of the current class and
// passing the current object as the constructor's argument.
// Performing all the actual copying in the constructor helps to
// keep the result consistent: the constructor will not return a
// result until the new object is fully built; thus, no object
// can have a reference to a partially-built clone.
class Rectangle extends Shape is
    field width: int
    field height: int

    constructor Rectangle(source: Rectangle) is
        // A parent constructor call is needed to copy private
        // fields defined in the parent class.
        super(source)
        this.width = source.width
        this.height = source.height

    method clone():Shape is
        return new Rectangle(this)


class Circle extends Shape is
    field radius: int

    constructor Circle(source: Circle) is
        super(source)
        this.radius = source.radius

    method clone():Shape is
        return new Circle(this)


// Somewhere in the client code.
class Application is
    field shapes: array of Shape

    constructor Application() is
        Circle circle = new Circle()
        circle.X = 10
        circle.Y = 10
        circle.radius = 20
        shapes.add(circle)

        Circle anotherCircle = circle.clone()
        shapes.add(anotherCircle)
        // The `anotherCircle` variable contains an exact copy
        // of the `circle` object.

        Rectangle rectangle = new Rectangle()
        rectangle.width = 10
        rectangle.height = 20
        shapes.add(rectangle)

    method businessLogic() is
        // Prototype rocks because it lets you produce a copy of
        // an object without knowing anything about its type.
        Array shapesCopy = new Array of Shapes.

        // For instance, we don't know the exact elements in the
        // shapes array. All we know is that they are all
        // shapes. But thanks to polymorphism, when we call the
        // `clone` method on a shape the program checks its real
        // class and runs the appropriate clone method defined
        // in that class. That's why we get proper clones
        // instead of a set of simple Shape objects.
        foreach (s in shapes) do
            shapesCopy.add(s.clone())

        // The `shapesCopy` array contains exact copies of the
        // `shape` array's children.
```

## Applicability

> **Use the Prototype pattern when your code shouldn’t depend on the concrete classes of objects that you need to copy.**

**Sử funjg prototype khi code của bạn không muốn phụ thuộc vào lớp cụ thể của đối tượng cần sao chép**

> This happens a lot when your code works with objects passed to you from 3rd-party code via some interface. The concrete classes of these objects are unknown, and you couldn’t depend on them even if you wanted to.

Chuyện này xảy ra nhiều, nếu code của bạn làm việc với các object được truyền từ bên thứ 3 thông qua một vài interface. Các lớp cụ thể của những object này là không xác định, nên bạn không thể phụ thuộc vào nó, ngay cả khi bạn muốn.

> The Prototype pattern provides the client code with a general interface for working with all objects that support cloning. This interface makes the client code independent from the concrete classes of objects that it clones.

Prototype pattern cung cấp cho client một interface chung để làm việc với tất cả các đối tượng hỗ trợ sao chép. Interface này giúp code của client độc lập với lớp cụ thể mà nó sao chép.

> **Use the pattern when you want to reduce the number of subclasses that only differ in the way they initialize their respective objects.**

**Sử dụng pattern prototype khi bạn muốn giảm số lượng lớp con, chỉ khác nhau về cách chúng khởi tạo đối tượng tương ứng.**

> Suppose you have a complex class that requires a laborious configuration before it can be used. There are several common ways to configure this class, and this code is scattered through your app. To reduce the duplication, you create several subclasses and put every common configuration code into their constructors. You solved the duplication problem, but now you have lots of dummy subclasses.

Giả sử bạn có một lớp phức tạp yêu cầu cấu hình tốn nhiều công sức trước khi nó có thể được sử dụng. Có một số cách phổ biến để cấu hình cho class này, và code này nằm rải rác trong ứng dụng của bạn. Để giảm sự trùng lặp code, bạn tạo một số lớp con và đặt code cấu hình  vào trong hàm khởi tạo. Bạn đã giải quyết được vấn đề trùng lặp code, nhưng bây giờ bạn có rất nhiều lớp con.

> The Prototype pattern lets you use a set of pre-built objects configured in various ways as prototypes. Instead of instantiating a subclass that matches some configuration, the client can simply look for an appropriate prototype and clone it.

Prototype pattern cho phép bạn sử dụng một tập các object được tạo sẵn, và được cấu hình theo nhiều cách khác nhau. Thay vì khởi tạo một lớp con phù hợp với cấu hình, client chỉ cần tìm kiếm một prototype thích hợp và clone nó.

## How to Implement

> 1. Create the prototype interface and declare the `clone` method in it. Or just add the method to all classes of an existing class hierarchy, if you have one.

1. Tạo interface prototype và định nghĩa hàm `clone`. Hoặc thêm phương thức clone vào tất cả các class.

> 2. A prototype class must define the alternative constructor that accepts an object of that class as an argument. The constructor must copy the values of all fields defined in the class from the passed object into the newly created instance. If you’re changing a subclass, you must call the parent constructor to let the superclass handle the cloning of its private fields.
If your programming language doesn’t support method overloading, you won’t be able to create a separate “prototype” constructor. Thus, copying the object’s data into the newly created clone will have to be performed within the `clone` method. Still, having this code in a regular constructor is safer because the resulting object is returned fully configured right after you call the `new` operator.

2. class prototype phải định nghĩa hàm khởi tạo thay thế để chấp nhận đối tượng của lớp đó như một tham số. hàm khởi tạo phải copy các giá trị của tất cả các biến được định nghĩa trong class từ object được truyền vào. Nếu bạn thay đổi một lớp con, bạn phải gọi đến hàm khởi tạo của lớp cha để cho phép lớp cha xử lý việc sao chép các biến private.

    Nếu ngôn ngữ đó không hỗ trợ overloading, bạn không thể tạo một hàm khởi tạo prototype. Do đó, việc sao chép data của object sang một object mới phải được thực hiện trong hàm `clone`. Tuy nhiên, đoạn code này trong hàm khởi tạo sẽ an toàn hơn vì đối tượng trả về được cấu hình đầy đủ ngay sau khi bạn gọi toán tử `new`.


> 3. The cloning method usually consists of just one line: running a `new` operator with the prototypical version of the constructor. Note, that every class must explicitly override the cloning method and use its own class name along with the `new` operator. Otherwise, the cloning method may produce an object of a parent class.

3. phương thức sao chép thường chỉ có một dòng. gọi toán tử `new` với hàm khởi tạo prototype. Chú ý rằng tất cả các class phải override phương thức clone, và sử dụng tên class của chính mình với toán tử `new`. Nếu không, phương thức clone có thể tạo ra một đối tượng của lớp cha.

> 4. Optionally, create a centralized prototype registry to store a catalog of frequently used prototypes.
You can implement the registry as a new factory class or put it in the base prototype class with a static method for fetching the prototype. This method should search for a prototype based on search criteria that the client code passes to the method. The criteria might either be a simple string tag or a complex set of search parameters. After the appropriate prototype is found, the registry should clone it and return the copy to the client.
Finally, replace the direct calls to the subclasses’ constructors with calls to the factory method of the prototype registry.

4. Tuỳ chọn, tạo một prototype registry để lưu trữ một tập các prototype hay sử dụng.

    Bạn có thể triển khai registry như là một class factory hoặc đặt nó vào một class prototype base với một static method để tìm prototype. Method này nên tìm kiếm prototype dựa trên tham số đầu vào mà client gửi đến. Điều kiện tìm kiếm có thể là một chuỗi tag hoặc một tập các tham số phức tạp. Sau khi tìm thấy prototype thích hợp, registry nên sao chép nó và trả về đối tượng đã sao chép cho client.
    Cuối cùng, thay thế các lệnh gọi trức tiếp đến hàm khởi tạo của lớp con bằng factory method của prototype registry

## Pros and Cons
### pros
>  - You can clone objects without coupling to their concrete classes.
>  - You can get rid of repeated initialization code in favor of cloning pre-built prototypes.
>  - You can produce complex objects more conveniently.
>  - You get an alternative to inheritance when dealing with configuration presets for complex objects.

- Bạn có thể clone object mà không cần quan tâm đến lớp cụ thể.
- Bạn có thể tránh lặp đi lặp lại việc tạo mới các object bằng việc sao chép các prototype có sẵn.
- Bạn có thể tạo một đối tượng phức tạp một cách thuận tiện.
- Bạn nhận được một giải pháp thay thế cho kế thừa khi phải xử lý các cấu hình cho các đối tượng phức tạp.

### Cons
>  - Cloning complex objects that have circular references might be very tricky.

- Sao chép các đối tượng phức tạp có tham chiếu vòng tròn có thể rất khó.

## Relations with Other Patterns

> - Many designs start by using **Factory Method** (less complicated and more customizable via subclasses) and evolve toward **Abstract Factory**, **Prototype**, or **Builder** (more flexible, but more complicated).

- Nhiều thiết kế bắt đầu bằng cách sử dụng **Factory Method**(ít phức tạp và dễ tuỳ chỉnh thông qua các lớp con) và phát triển theo hướng **Abstract Factory**, **Prototype** hoặc **Builder**(linh hoạt hơn, nhưng phức tạp).

> - **Abstract Factory** classes are often based on a set of **Factory Methods**, but you can also use **Prototype** to compose the methods on these classes.

- các lớp **Abstract Factory** thường base trên một tập các **Factory Method**, nhưng bạn cũng có thể sử dụng **prototype** để cấu trúc các phương thức trong các lớp này.

> - **Prototype** can help when you need to save copies of **Commands** into history.

- **Prototype** có thể hữu ích khi bạn cần lưu các bản sao của **Commands** vào lịch sử.

> - Designs that make heavy use of **Composite** and **Decorator** can often benefit from using **Prototype**. Applying the pattern lets you clone complex structures instead of re-constructing them from scratch.

- Các thiết kế sử dụng nhiều **Composite** và **Decorator** thường có thể được hưởng lợi từ việc sử dụng **Prototype**. Áp dụng pattern cho phép bạn sao chép các cấu trúc phức tạp thay vì xây dựng lại chúng từ đầu.

> - **Prototype** isn’t based on inheritance, so it doesn’t have its drawbacks. On the other hand, Prototype requires a complicated initialization of the cloned object. **Factory Method** is based on inheritance but doesn’t require an initialization step.

- **Prototype** không dựa trên sự kế thừa, vì vậy nó không có nhược điểm. Mặt khác, Prototype yêu cầu khởi tạo nhân bản đối tượng phức tạp. **Factory Method** dựa trên kế thừa nhưng không yêu cầu bước khởi tạo.

> - Sometimes **Prototype** can be a simpler alternative to **Memento**. This works if the object, the state of which you want to store in the history, is fairly straightforward and doesn’t have links to external resources, or the links are easy to re-establish.

- Đôi khi **Prototype** có thể là một giải pháp thay thế đơn giản hơn cho **Memento**. Điều này hoạt động nếu đối tượng, trạng thái mà bạn muốn lưu trữ trong lịch sử, khá đơn giản và không có liên kết đến tài nguyên bên ngoài hoặc các liên kết dễ thiết lập lại.

> - **Abstract Factories**, **Builders** and **Prototypes** can all be implemented as **Singletons**.

- Tất cả các **Abstract Factory**, **Builder** và **Prototype** đều có thể được triển khai dưới dạng các **Singleton**.
