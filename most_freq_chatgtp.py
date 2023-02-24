#Esto es ChatGPT
def most_frequent(arr):
    n = len(arr)
    max_count = 0
    most_freq = None
    left, right = 0, n-1
    while left < right:
        mid = (left + right) // 2
        if arr[mid] == arr[mid-1]:
            count = 1
            i = mid - 1
            while i >= left and arr[i] == arr[mid]:
                count += 1
                i -= 1
            i = mid + 1
            while i <= right and arr[i] == arr[mid]:
                count += 1
                i += 1
            if count > max_count:
                max_count = count
                most_freq = arr[mid]
        elif arr[mid] == arr[mid+1]:
            count = 1
            i = mid + 1
            while i <= right and arr[i] == arr[mid]:
                count += 1
                i += 1
            i = mid - 1
            while i >= left and arr[i] == arr[mid]:
                count += 1
                i -= 1
            if count > max_count:
                max_count = count
                most_freq = arr[mid]
        else:
            return most_freq
    return most_freq

arr = [1,2,2,3,3,3,4,4]

print(most_frequent(arr))

