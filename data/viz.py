import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt

# Load data
data = pd.read_csv("2024_09_28_Leipzig_Golang_Meetup_Total_and_Active_Members.csv", sep=",")
data.columns = ["date", "total", "active"]
data["date"] = pd.to_datetime(data.date)

sns.set_theme(rc={'figure.figsize':(11.7,8.27)})
sns.set_style("darkgrid")

# Create a line plot using Seaborn
sns.set_theme(style="whitegrid")
sns.lineplot(x="date", y="active", data=data)
sns.lineplot(x="date", y="total", data=data)

# Add title and labels
plt.title("golangleipzig.space: active and total members (2021-2024)")
plt.xlabel("date")
plt.ylabel("members")

# Rotate x-axis labels for better readability
plt.xticks(rotation=45)

# Save the plot to a file
plt.savefig("active_cases_plot.png")
