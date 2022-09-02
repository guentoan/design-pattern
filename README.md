# Design patterns

## Nguồn
Lúc đầu định ngồi viết code và tổng hợp các design patterns, nhưng sau tìm được [nguồn này](https://refactoring.guru/design-patterns) có cả giải thích và code, nên ngồi dịch luôn.

## Mục lục

 - [Creational Design Patterns](/creational-patterns/README.md)
   - [Factory Method](/creational-patterns/factory-method/README.md)
   - [Abstract Factory](/creational-patterns/abstract-factory/README.md)
   - [Builder](/creational-patterns/builder/README.md)
   - [Prototype](/creational-patterns/prototype/README.md)
   - [Singleton](/creational-patterns/singleton/README.md)
 - [Structural Design Patterns](/structural-patterns/README.md)
   - [Adapter](/structural-patterns/adapter/README.md)
   - [Bridge](/structural-patterns/bridge/README.md)
   - [Composite](/structural-patterns/composite/README.md)
   - [Decorator](/structural-patterns/decorator/README.md)
   - [Facade](/structural-patterns/facade/README.md)
   - [Flyweight](/structural-patterns/flyweight/README.md)
   - [Proxy](/structural-patterns/proxy/README.md)
 - [Behavioral Design Patterns](/behavioral-patterns/README.md)
   - [Chain of responsibility](/behavioral-patterns/chain-of-responsibility/README.md)
   - [Command](/behavioral-patterns/command/README.md)
   - [Iterator](/behavioral-patterns/iterator/README.md)
   - [Mediator](/behavioral-patterns/mediator/README.md)
   - [Memento](/behavioral-patterns/memento/README.md)
   - [Observer](/behavioral-patterns/observer/README.md)
   - [State](/behavioral-patterns/state/README.md)
   - [Strategy](/behavioral-patterns/strategy/README.md)
   - [Template method](/behavioral-patterns/template-method/README.md)
   - [Visitor](/behavioral-patterns/visitor/README.md)

## What's a design pattern?

>Design patterns are typical solutions to commonly occurring problems in software design. They are like pre-made blueprints that you can customize to solve a recurring design problem in your code.

Design Pattern là một giải pháp tổng thể cho các vấn đề chung trong thiết kế phần mềm. Nó cũng tương tự các bản vẽ kỹ thuật, mà ở đó chúng được dùng để giải quyết các vấn đề lặp đi lặp lại trong code của bạn.

>You can’t just find a pattern and copy it into your program, the way you can with off-the-shelf functions or libraries. The pattern is not a specific piece of code, but a general concept for solving a particular problem. You can follow the pattern details and implement a solution that suits the realities of your own program.

Các design pattern không thể copy rồi paste như cách bạn làm với các function có sẵn hay thư viện. Các pattern không phải là những đoạn code cụ thể, mà nó là những khái niệm tổng quát để giải quyết các vấn đề riêng biệt. Bạn có thể tìm hiểu các design pattern và triển khai chúng lên ứng dụng của bạn.

> Patterns are often confused with algorithms, because both concepts describe typical solutions to some known problems. While an algorithm always defines a clear set of actions that can achieve some goal, a pattern is a more high-level description of a solution. The code of the same pattern applied to two different programs may be different.

Các pattern thường bị nhầm lẫn với thuật toán, vì chúng đều là những khái niệm mô tả giải pháp cho một vấn đề nào đó. Trong khi thuật toán là định nghĩa những hành động cụ thể để giải quyết vấn đề thì design pattern lại là một mô tả ở mức cao hơn cho các giải pháp. Code của một pattern có thể khác nhau nếu chúng được triển khai trên hai ứng dụng khác nhau.

>An analogy to an algorithm is a cooking recipe: both have clear steps to achieve a goal. On the other hand, a pattern is more like a blueprint: you can see what the result and its features are, but the exact order of implementation is up to you.

Một thuật toán cũng tương tự như công thức nấu ăn: cả hai đều có các bước rõ ràng để giải quyết vấn đề. Mặt khác, Pattern giống như một bản vẽ kỹ thuật: ở đó bạn có thể thấy kết quả và tính năng của chúng, nhưng thứ tự thực hiện các bước lại tuỳ thuộc vào bạn.

## What does the pattern consist of?

>Most patterns are described very formally so people can reproduce them in many contexts. Here are the sections that are usually present in a pattern description:

Hầu hết các tài liệu mô tả khá chi tiết, và bạn có thể tái sử dụng cho nhiều trường hợp. Dưới đây là các thành phần thường có trong xuyên suốt cuốn sách:

> - **Intent** of the pattern briefly describes both the problem and the solution.
 - **Mục đích** của pattern: mô tả ngắn gọn vấn đề và giải pháp.

> - **Motivation** further explains the problem and the solution the pattern makes possible.
 - **Motivation**: giải thích thêm vấn đề và giải pháp với pattern khả thi.

> - **Structure** of classes shows each part of the pattern and how they are related.
- **Cấu trúc** của các lớp cho thấy từng phần của pattern và cách chúng liên kết với nhau.

> - **Code example** in one of the popular programming languages makes it easier to grasp the idea behind the pattern.
 - **Ví dụ** bằng những ngôn ngữ lập trình phổ biến để giúp bạn dễ ràng nắm bắt ý tường của pattern.


>Some pattern catalogs list other useful details, such as applicability of the pattern, implementation steps and relations with other patterns.

Một số pattern catalogs liệt kê các cách sử dụng hữu ích như: tính ứng dụng của các pattern, các bước triển khai và mối liên hệ với các patterns khác.

## Why should I learn patterns?

> The truth is that you might manage to work as a programmer for many years without knowing about a single pattern. A lot of people do just that. Even in that case, though, you might be implementing some patterns without even knowing it. So why would you spend time learning them?

Sự thật là các lập trình viên có thể xoay xở làm việc trong nhiều năm mà không cần biết đến bất kỳ pattern nào. Rất nhiều người làm như vậy. Tuy nhiên, ngay cả trong trường hợp đó, bạn có thể đang triển khai một số pattern mà không hề hay biết. Vậy tại sao bạn lại dành thời gian tìm hiểu chúng?

> - Design patterns are a toolkit of tried and tested solutions to common problems in software design. Even if you never encounter these problems, knowing patterns is still useful because it teaches you how to solve all sorts of problems using principles of object-oriented design.

 - Design patterns là một công cụ bao gồm các giải pháp để giải quyết các vấn đề thường gặp trong thiết kế phần mềm. Ngay cả khi bạn không bao giờ gặp các vấn đề này, thì việc biết các patterns vẫn hữu ích, vì nó dạy bạn cách giải quyết tất cả các vấn đề bằng cách sử dụng nguyên tắc của thiết kế hương đối tượng.

 > - Design patterns define a common language that you and your teammates can use to communicate more efficiently. You can say, “Oh, just use a Singleton for that,” and everyone will understand the idea behind your suggestion. No need to explain what a singleton is if you know the pattern and its name.

 - Design patterns được tạo bởi một ngôn ngữ chung, ở đó bạn và team bạn có thể sử dụng để giao tiếp hiệu quả hơn. Bạn chỉ cần nói:"Chỗ này dùng singleton nhé" và tất cả mọi người có thể hiểu ý tưởng của bạn. Không cần phải giải thích singleton là gì nếu bạn biết pattern đó.

