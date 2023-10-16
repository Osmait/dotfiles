

from typing import Generic, TypeVar


T = TypeVar("T")


class StackOverflowError(BaseException):
    pass


class StackUnderflowError(BaseException):
    pass


class Stack(Generic[T]):
    """A stack is an abstract data type that serves as a collection of
    elements with two principal operations: push() and pop(). push() adds an
    element to the top of the stack, and pop() removes an element from the top
    of a stack. The order in which elements come off of a stack are
    Last In, First Out (LIFO).
    https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
    """

    def __init__(self, limit: int = 10) -> None:
        self.stack: list[T] = []
        self.limit = limit

    def __bool__(self) -> bool:
        return bool(self.stack)

    def __str__(self) -> str:
        return str(self.stack)

    def push(self, data: T) -> None:
        """Push an element to the top of the stack."""
        if len(self.stack) >= self.limit:
            raise StackOverflowError
        self.stack.append(data)

    def pop(self) -> T:
        """
        Pop an element off of the top of the stack.

        >>> Stack().pop()
        Traceback (most recent call last):
            ...
        data_structures.stacks.stack.StackUnderflowError
        """
        if not self.stack:
            raise StackUnderflowError
        return self.stack.pop()

    def peek(self) -> T:
        """
        Peek at the top-most element of the stack.

        >>> Stack().pop()
        Traceback (most recent call last):
            ...
        data_structures.stacks.stack.StackUnderflowError
        """
        if not self.stack:
            raise StackOverflowError
        return self.stack[-1]
