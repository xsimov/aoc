defmodule Day2 do
  def find_wc_pin(instructions) do
    execute_lines(instructions, State.start(5))
  end

  defp execute_lines([first_line | other_lines], pid) do
    result = execute_commands(first_line, pid)
    [result | execute_lines(other_lines, pid)]
  end
  defp execute_lines([], _), do: []

  defp execute_commands([first | others], pid) do
    process_order(first, State.get(pid))
    |> State.set(pid)
    execute_commands(others, pid)
  end
  defp execute_commands([], pid), do: State.get(pid)

  defp process_order("U", num), do: go_up(num)
  defp process_order("R", num), do: go_right(num)
  defp process_order("L", num), do: go_left(num)
  defp process_order("D", num), do: go_down(num)

  defp go_up(num) when num == 3, do: 1
  defp go_up(num) when num in 6..8, do: num - 4
  defp go_up(num) when num == "A", do: 6
  defp go_up(num) when num == "B", do: 7
  defp go_up(num) when num == "C", do: 8
  defp go_up(num) when num == "D", do: "B"
  defp go_up(num), do: num

  defp go_down(num) when num == 1, do: 3
  defp go_down(num) when num in 2..4, do: num + 4
  defp go_down(num) when num == 6, do: "A"
  defp go_down(num) when num == 7, do: "B"
  defp go_down(num) when num == 8, do: "C"
  defp go_down(num) when num == "B", do: "D"
  defp go_down(num), do: num

  defp go_right(num) when num in [2, 3, 5, 6, 7, 8], do: num + 1
  defp go_right(num) when num == "A", do: "B"
  defp go_right(num) when num == "B", do: "C"
  defp go_right(num), do: num

  defp go_left(num) when num in [3, 4, 6, 7, 8, 9], do: num - 1
  defp go_left(num) when num == "B", do: "A"
  defp go_left(num) when num == "C", do: "B"
  defp go_left(num), do: num
end

defmodule State do
  def start(initial_state) do
    spawn(fn -> loop(initial_state) end)
  end

  defp loop(state) do
    receive do
      {:set, new_state} -> loop(new_state)
      {:get, caller} ->
        send caller, {:got, state}
        loop(state)
    end
  end

  def get(pid) do
    send pid, {:get, self()}
    receive do
      {:got, state} -> state
    end
  end

  def set(state, pid) do
    send pid, {:set, state}
  end
end
